package connmgr

import (
	"time"
)

func DecayNone() DecayFn { _ = "STUB: not implemented"; return *new(DecayFn) }

func DecayFixed(minuend int) DecayFn { _ = "STUB: not implemented"; return *new(DecayFn) }

func DecayLinear(coef float64) DecayFn { _ = "STUB: not implemented"; return *new(DecayFn) }

func DecayExpireWhenInactive(after time.Duration) DecayFn {
	_ = "STUB: not implemented"
	return *new(DecayFn)
}

func BumpSumUnbounded() BumpFn { _ = "STUB: not implemented"; return *new(BumpFn) }

func BumpSumBounded(min, max int) BumpFn { _ = "STUB: not implemented"; return *new(BumpFn) }

func BumpOverwrite() BumpFn { _ = "STUB: not implemented"; return *new(BumpFn) }
