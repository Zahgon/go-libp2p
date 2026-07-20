//go:build !windows && !riscv64 && !loong64

package tcp

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/transport"
	"github.com/marten-seemann/tcp"
	"github.com/mikioh/tcpinfo"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	newConns      *prometheus.CounterVec
	closedConns   *prometheus.CounterVec
	segsSentDesc  *prometheus.Desc
	segsRcvdDesc  *prometheus.Desc
	bytesSentDesc *prometheus.Desc
	bytesRcvdDesc *prometheus.Desc
)

const collectFrequency = 10 * time.Second

var defaultCollector *aggregatingCollector

var initMetricsOnce sync.Once

func initMetrics() { _ = "STUB: not implemented"; return }

type aggregatingCollector struct {
	cronOnce sync.Once

	mutex                sync.Mutex
	highestID            uint64
	conns                map[uint64]*tracingConn
	rtts                 prometheus.Histogram
	connDurations        prometheus.Histogram
	segsSent, segsRcvd   uint64
	bytesSent, bytesRcvd uint64
}

var _ prometheus.Collector = &aggregatingCollector{}

func newAggregatingCollector() *aggregatingCollector { _ = "STUB: not implemented"; return nil }

func (c *aggregatingCollector) AddConn(t *tracingConn) uint64 { _ = "STUB: not implemented"; return 0 }

func (c *aggregatingCollector) removeConn(id uint64) { _ = "STUB: not implemented"; return }

func (c *aggregatingCollector) Describe(descs chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

func (c *aggregatingCollector) cron() { _ = "STUB: not implemented"; return }

func (c *aggregatingCollector) gatherMetrics(now time.Time) { _ = "STUB: not implemented"; return }

func (c *aggregatingCollector) Collect(metrics chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

func (c *aggregatingCollector) ClosedConn(conn *tracingConn, direction string) {
	_ = "STUB: not implemented"
	return
}

type tracingConn struct {
	id uint64

	collector *aggregatingCollector

	startTime time.Time
	isClient  bool

	manet.Conn
	tcpConn   *tcp.Conn
	closeOnce sync.Once
	closeErr  error
}

func newTracingConn(c manet.Conn, collector *aggregatingCollector, isClient bool) (*tracingConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *tracingConn) getDirection() string { _ = "STUB: not implemented"; return "" }

func (c *tracingConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *tracingConn) getTCPInfo() (*tcpinfo.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tracingListener struct {
	transport.GatedMaListener
	collector *aggregatingCollector
}

func newTracingListener(l transport.GatedMaListener, collector *aggregatingCollector) *tracingListener {
	_ = "STUB: not implemented"
	return nil
}

func (l *tracingListener) Accept() (manet.Conn, network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), *new(network.ConnManagementScope), nil
}
