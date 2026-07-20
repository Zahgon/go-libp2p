//go:build !plan9

package reuseport

func reuseErrShouldRetry(err error) bool { _ = "STUB: not implemented"; return false }
