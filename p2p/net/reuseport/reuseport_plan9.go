package reuseport

const (
	EADDRINUSE   = "address in use"
	ECONNREFUSED = "connection refused"
)

func reuseErrShouldRetry(err error) bool { _ = "STUB: not implemented"; return false }
