package tcpreuse

import (
	"os"
	"strings"
)

const envReuseport = "LIBP2P_TCP_REUSEPORT"

var EnvReuseportVal = true

func init() {
	v := strings.ToLower(os.Getenv(envReuseport))
	if v == "false" || v == "f" || v == "0" {
		EnvReuseportVal = false
		log.Info("REUSEPORT disabled", "LIBP2P_TCP_REUSEPORT", v)
	}
}

func ReuseportIsAvailable() bool { _ = "STUB: not implemented"; return false }
