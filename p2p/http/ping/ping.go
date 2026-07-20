package httpping

import (
	"net/http"
)

const pingSize = 32
const PingProtocolID = "/http-ping/1"

type Ping struct{}

var _ http.Handler = Ping{}

func (Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func SendPing(client http.Client) error { _ = "STUB: not implemented"; return nil }
