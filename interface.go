package mapfactory

type Map[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (value V, ok bool)
	Len() int
	Range(func(key K, value V) bool)
}
