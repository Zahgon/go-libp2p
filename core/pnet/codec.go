package pnet

import (
	"bufio"
	"io"
)

var (
	pathPSKv1  = []byte("/key/swarm/psk/1.0.0/")
	pathBin    = "/bin/"
	pathBase16 = "/base16/"
	pathBase64 = "/base64/"
)

func readHeader(r *bufio.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func expectHeader(r *bufio.Reader, expected []byte) error { _ = "STUB: not implemented"; return nil }

func DecodeV1PSK(in io.Reader) (PSK, error) { _ = "STUB: not implemented"; return *new(PSK), nil }
