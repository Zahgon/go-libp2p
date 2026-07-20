package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

const Protocol = "/proxy-example/0.0.1"

func makeRandomHost(port int) host.Host { _ = "STUB: not implemented"; return *new(host.Host) }

type ProxyService struct {
	host      host.Host
	dest      peer.ID
	proxyAddr ma.Multiaddr
}

func NewProxyService(h host.Host, proxyAddr ma.Multiaddr, dest peer.ID) *ProxyService {
	_ = "STUB: not implemented"
	return nil
}

func streamHandler(stream network.Stream) { _ = "STUB: not implemented"; return }

func (p *ProxyService) Serve() { _ = "STUB: not implemented"; return }

func (p *ProxyService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func addAddrToPeerstore(h host.Host, addr string) peer.ID {
	_ = "STUB: not implemented"
	return *new(peer.ID)
}

const help = `
This example creates a simple HTTP Proxy using two libp2p peers. The first peer
provides an HTTP server locally which tunnels the HTTP requests with libp2p
to a remote peer. The remote peer performs the requests and 
send the sends the response back.

Usage: Start remote peer first with:   ./proxy
       Then start the local peer with: ./proxy -d <remote-peer-multiaddress>

Then you can do something like: curl -x "localhost:9900" "http://ipfs.io".
This proxies sends the request through the local peer, which proxies it to
the remote peer, which makes it and sends the response back.`

func main() {
	flag.Usage = func() {
		fmt.Println(help)
		flag.PrintDefaults()
	}

	destPeer := flag.String("d", "", "destination peer address")
	port := flag.Int("p", 9900, "proxy port")
	p2pport := flag.Int("l", 12000, "libp2p listen port")
	flag.Parse()

	if *destPeer != "" {

		host := makeRandomHost(*p2pport + 1)

		destPeerID := addAddrToPeerstore(host, *destPeer)
		proxyAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", *port))
		if err != nil {
			log.Fatalln(err)
		}

		proxy := NewProxyService(host, proxyAddr, destPeerID)
		proxy.Serve()
	} else {
		host := makeRandomHost(*p2pport)

		_ = NewProxyService(host, nil, "")
		<-make(chan struct{})
	}

}
