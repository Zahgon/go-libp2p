package metricshelper

import (
	"sync"
)

const capacity = 8

var stringPool = sync.Pool{New: func() any {
	s := make([]string, 0, capacity)
	return &s
}}

func GetStringSlice() *[]string { _ = "STUB: not implemented"; return nil }

func PutStringSlice(s *[]string) { _ = "STUB: not implemented"; return }
