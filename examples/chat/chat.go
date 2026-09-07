package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"os"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
)

func handleStream(s network.Stream) { _ = "STUB: not implemented"; return }

func readData(rw *bufio.ReadWriter) { _ = "STUB: not implemented"; return }

func writeData(rw *bufio.ReadWriter) { _ = "STUB: not implemented"; return }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sourcePort := flag.Int("sp", 0, "Source port number")
	dest := flag.String("d", "", "Destination multiaddr string")
	help := flag.Bool("help", false, "Display help")
	debug := flag.Bool("debug", false, "Debug generates the same node ID on every execution")

	flag.Parse()

	if *help {
		fmt.Printf("This program demonstrates a simple p2p chat application using libp2p\n\n")
		fmt.Println("Usage: Run './chat -sp <SOURCE_PORT>' where <SOURCE_PORT> can be any port number.")
		fmt.Println("Now run './chat -d <MULTIADDR>' where <MULTIADDR> is multiaddress of previous listener host.")

		os.Exit(0)
	}

	var r io.Reader
	if *debug {

		r = mrand.New(mrand.NewSource(int64(*sourcePort)))
	} else {
		r = rand.Reader
	}

	h, err := makeHost(*sourcePort, r)
	if err != nil {
		log.Println(err)
		return
	}

	if *dest == "" {
		startPeer(ctx, h, handleStream)
	} else {
		rw, err := startPeerAndConnect(ctx, h, *dest)
		if err != nil {
			log.Println(err)
			return
		}

		go writeData(rw)
		go readData(rw)

	}

	select {}
}

func makeHost(port int, randomness io.Reader) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func startPeer(_ context.Context, h host.Host, streamHandler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func startPeerAndConnect(_ context.Context, h host.Host, destination string) (*bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
