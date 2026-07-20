package pnet

import "os"

const EnvKey = "LIBP2P_FORCE_PNET"

var ForcePrivateNetwork = false

func init() {
	ForcePrivateNetwork = os.Getenv(EnvKey) == "1"
}
