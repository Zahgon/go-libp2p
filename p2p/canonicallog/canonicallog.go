package canonicallog

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/multiformats/go-multiaddr"
)

var log = slog.New(
	slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{
			Level:     logging.ConfigFromEnv().LevelForSystem("canonical-log"),
			AddSource: true}))

func logWithSkip(ctx context.Context, l *slog.Logger, level slog.Level, skip int, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func LogMisbehavingPeer(p peer.ID, peerAddr multiaddr.Multiaddr, component string, err error, msg string) {
	_ = "STUB: not implemented"
	return
}

func LogMisbehavingPeerNetAddr(p peer.ID, peerAddr net.Addr, component string, originalErr error, msg string) {
	_ = "STUB: not implemented"
	return
}

func LogPeerStatus(sampleRate int, p peer.ID, peerAddr multiaddr.Multiaddr, keyVals ...string) {
	_ = "STUB: not implemented"
	return
}
