package pstoreds

type cache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Add(key K, value V)
	Remove(key K)
	Contains(key K) bool
	Peek(key K) (value V, ok bool)
	Keys() []K
}

type noopCache[K comparable, V any] struct {
}

var _ cache[int, int] = (*noopCache[int, int])(nil)

func (*noopCache[K, V]) Get(_ K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (*noopCache[K, V]) Add(_ K, _ V) { _ = "STUB: not implemented"; return }

func (*noopCache[K, V]) Remove(_ K) { _ = "STUB: not implemented"; return }

func (*noopCache[K, V]) Contains(_ K) bool { _ = "STUB: not implemented"; return false }

func (*noopCache[K, V]) Peek(_ K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (*noopCache[K, V]) Keys() (keys []K) { _ = "STUB: not implemented"; return nil }
