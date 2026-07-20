package main

import (
	"context"
	"flag"
	"io"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	golog "github.com/ipfs/go-log/v2"
)

func makeRoutedHost(listenPort int, randseed int64, bootstrapPeers []peer.AddrInfo, globalFlag string) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func main() {

	golog.SetAllLoggers(golog.LevelInfo)

	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	global := flag.Bool("global", false, "use global ipfs peers for bootstrapping")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	var bootstrapPeers []peer.AddrInfo
	var globalFlag string
	if *global {
		log.Println("using global bootstrap")
		bootstrapPeers = IPFS_PEERS
		globalFlag = " -global"
	} else {
		log.Println("using local bootstrap")
		bootstrapPeers = getLocalPeerInfo()
		globalFlag = ""
	}
	ha, err := makeRoutedHost(*listenF, *seed, bootstrapPeers, globalFlag)
	if err != nil {
		log.Fatal(err)
	}

	ha.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
		log.Println("Got a new stream!")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	if *target == "" {
		log.Println("listening for connections")
		select {}
	}

	peerid, err := peer.Decode(*target)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("opening stream")

	s, err := ha.NewStream(context.Background(), peerid, "/echo/1.0.0")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = s.Write([]byte("Hello, world!\n"))
	if err != nil {
		log.Fatalln(err)
	}

	out, err := io.ReadAll(s)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("read reply: %q\n", out)
}

func doEcho(s network.Stream) error { _ = "STUB: not implemented"; return nil }
