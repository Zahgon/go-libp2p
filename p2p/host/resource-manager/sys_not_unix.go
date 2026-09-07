//go:build !linux && !darwin && !windows

package rcmgr

func getNumFDs() int { _ = "STUB: not implemented"; return 0 }
