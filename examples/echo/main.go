package main

import (
	"context"
	"flag"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"

	golog "github.com/ipfs/go-log/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	golog.SetAllLoggers(golog.LevelInfo)

	listenF := flag.Int("l", 0, "wait for incoming connections")
	targetF := flag.String("d", "", "target peer to dial")
	insecureF := flag.Bool("insecure", false, "use an unencrypted connection")
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	ha, err := makeBasicHost(*listenF, *insecureF, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	if *targetF == "" {
		startListener(ctx, ha, *listenF, *insecureF)

		<-ctx.Done()
	} else {
		runSender(ctx, ha, *targetF)
	}
}

func makeBasicHost(listenPort int, insecure bool, randseed int64) (host.Host, error) {
	_ = "STUB: not implemented"
	return *new(host.Host), nil
}

func getHostAddress(ha host.Host) string { _ = "STUB: not implemented"; return "" }

func startListener(_ context.Context, ha host.Host, listenPort int, insecure bool) {
	_ = "STUB: not implemented"
	return
}

func runSender(_ context.Context, ha host.Host, targetPeer string) {
	_ = "STUB: not implemented"
	return
}

func doEcho(s network.Stream) error { _ = "STUB: not implemented"; return nil }
