package obs

import (
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
)

var MustRegisterWith = rcmgr.MustRegisterWith

type StatsTraceReporter = rcmgr.StatsTraceReporter

var NewStatsTraceReporter = rcmgr.NewStatsTraceReporter
