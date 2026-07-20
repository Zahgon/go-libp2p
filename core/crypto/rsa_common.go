package crypto

import (
	"fmt"
	"os"
)

const WeakRsaKeyEnv = "LIBP2P_ALLOW_WEAK_RSA_KEYS"

var MinRsaKeyBits = 2048

var maxRsaKeyBits = 8192

var ErrRsaKeyTooSmall error
var ErrRsaKeyTooBig error = fmt.Errorf("rsa keys must be <= %d bits", maxRsaKeyBits)

func init() {
	if _, ok := os.LookupEnv(WeakRsaKeyEnv); ok {
		MinRsaKeyBits = 512
	}

	ErrRsaKeyTooSmall = fmt.Errorf("rsa keys must be >= %d bits to be useful", MinRsaKeyBits)
}
