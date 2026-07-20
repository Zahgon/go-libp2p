package main

type config struct {
	RendezvousString string
	ProtocolID       string
	listenHost       string
	listenPort       int
}

func parseFlags() *config { _ = "STUB: not implemented"; return nil }
