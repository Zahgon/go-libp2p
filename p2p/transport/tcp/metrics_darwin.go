//go:build darwin

package tcp

import "github.com/mikioh/tcpinfo"

const (
	hasSegmentCounter = true
	hasByteCounter    = true
)

func getSegmentsSent(info *tcpinfo.Info) uint64 { _ = "STUB: not implemented"; return 0 }
func getSegmentsRcvd(info *tcpinfo.Info) uint64 { _ = "STUB: not implemented"; return 0 }
func getBytesSent(info *tcpinfo.Info) uint64    { _ = "STUB: not implemented"; return 0 }
func getBytesRcvd(info *tcpinfo.Info) uint64    { _ = "STUB: not implemented"; return 0 }
