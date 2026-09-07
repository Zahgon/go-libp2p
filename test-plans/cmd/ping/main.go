package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	libp2pwebrtc "github.com/libp2p/go-libp2p/p2p/transport/webrtc"

	"github.com/go-redis/redis/v8"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	libp2pquic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
	libp2pwebtransport "github.com/libp2p/go-libp2p/p2p/transport/webtransport"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func main() {
	var (
		transport      = os.Getenv("transport")
		muxer          = os.Getenv("muxer")
		secureChannel  = os.Getenv("security")
		isDialerStr    = os.Getenv("is_dialer")
		ip             = os.Getenv("ip")
		redisAddr      = os.Getenv("redis_addr")
		testTimeoutStr = os.Getenv("test_timeout_seconds")
	)

	testTimeout := 3 * time.Minute
	if testTimeoutStr != "" {
		secs, err := strconv.ParseInt(testTimeoutStr, 10, 32)
		if err == nil {
			testTimeout = time.Duration(secs) * time.Second
		}
	}

	if ip == "" {
		ip = "0.0.0.0"
	}

	if redisAddr == "" {
		redisAddr = "redis:6379"
	}

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	rClient := redis.NewClient(&redis.Options{
		DialTimeout: testTimeout,
		Addr:        redisAddr,
		Password:    "",
		DB:          0,
	})
	defer rClient.Close()

	for {
		if ctx.Err() != nil {
			log.Fatal("timeout waiting for redis")
		}

		_, err := rClient.Ping(ctx).Result()
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	isDialer := isDialerStr == "true"

	options := []libp2p.Option{}

	var listenAddr string
	switch transport {
	case "ws":
		options = append(options, libp2p.Transport(websocket.New))
		listenAddr = fmt.Sprintf("/ip4/%s/tcp/0/ws", ip)
	case "wss":
		options = append(options, libp2p.Transport(websocket.New, websocket.WithTLSConfig(generateTLSConfig()), websocket.WithTLSClientConfig(&tls.Config{InsecureSkipVerify: true})))
		listenAddr = fmt.Sprintf("/ip4/%s/tcp/0/wss", ip)
	case "tcp":
		options = append(options, libp2p.Transport(tcp.NewTCPTransport))
		listenAddr = fmt.Sprintf("/ip4/%s/tcp/0", ip)
	case "quic-v1":
		options = append(options, libp2p.Transport(libp2pquic.NewTransport))
		listenAddr = fmt.Sprintf("/ip4/%s/udp/0/quic-v1", ip)
	case "webtransport":
		options = append(options, libp2p.Transport(libp2pwebtransport.New))
		listenAddr = fmt.Sprintf("/ip4/%s/udp/0/quic-v1/webtransport", ip)
	case "webrtc-direct":
		options = append(options, libp2p.Transport(libp2pwebrtc.New))
		listenAddr = fmt.Sprintf("/ip4/%s/udp/0/webrtc-direct", ip)
	default:
		log.Fatalf("Unsupported transport: %s", transport)
	}
	options = append(options, libp2p.ListenAddrStrings(listenAddr))

	var skipMuxer bool
	var skipSecureChannel bool
	switch transport {
	case "quic-v1":
		fallthrough
	case "webtransport":
		fallthrough
	case "webrtc-direct":
		skipMuxer = true
		skipSecureChannel = true
	}

	if !skipSecureChannel {
		switch secureChannel {
		case "tls":
			options = append(options, libp2p.Security(libp2ptls.ID, libp2ptls.New))
		case "noise":
			options = append(options, libp2p.Security(noise.ID, noise.New))
		default:
			log.Fatalf("Unsupported secure channel: %s", secureChannel)
		}
	}

	if !skipMuxer {
		switch muxer {
		case "yamux":
			options = append(options, libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport))
		default:
			log.Fatalf("Unsupported muxer: %s", muxer)
		}
	}

	host, err := libp2p.New(options...)

	if err != nil {
		log.Fatalf("failed to instantiate libp2p instance: %s", err)
	}
	defer host.Close()

	log.Println("My multiaddr is: ", host.Addrs())

	if isDialer {
		val, err := rClient.BLPop(ctx, testTimeout, "listenerAddr").Result()
		if err != nil {
			log.Fatal("Failed to wait for listener to be ready")
		}
		otherMa := ma.StringCast(val[1])
		log.Println("Other peer multiaddr is: ", otherMa)
		otherMa, p2pComponent := ma.SplitLast(otherMa)
		otherPeerId, err := peer.Decode(p2pComponent.Value())
		if err != nil {
			log.Fatal("Failed to get peer id from multiaddr")
		}

		handshakeStartInstant := time.Now()
		err = host.Connect(ctx, peer.AddrInfo{
			ID:    otherPeerId,
			Addrs: []ma.Multiaddr{otherMa},
		})
		if err != nil {
			log.Fatal("Failed to connect to other peer")
		}

		ping := ping.NewPingService(host)

		res := <-ping.Ping(ctx, otherPeerId)
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		handshakePlusOneRTT := time.Since(handshakeStartInstant)

		testResult := struct {
			HandshakePlusOneRTTMillis float32 `json:"handshakePlusOneRTTMillis"`
			PingRTTMilllis            float32 `json:"pingRTTMilllis"`
		}{
			HandshakePlusOneRTTMillis: float32(handshakePlusOneRTT.Microseconds()) / 1000,
			PingRTTMilllis:            float32(res.RTT.Microseconds()) / 1000,
		}

		testResultJSON, err := json.Marshal(testResult)
		if err != nil {
			log.Fatalf("Failed to marshal test result: %v", err)
		}
		fmt.Println(string(testResultJSON))
	} else {
		var listenAddr ma.Multiaddr
		for _, addr := range host.Addrs() {
			if !manet.IsIPLoopback(addr) {
				listenAddr = addr
				break
			}
		}
		_, err := rClient.RPush(ctx, "listenerAddr", listenAddr.Encapsulate(ma.StringCast("/p2p/"+host.ID().String())).String()).Result()
		if err != nil {
			log.Fatal("Failed to send listener address")
		}
		time.Sleep(testTimeout)
		os.Exit(1)
	}
}

func generateTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }
