package iterator

type (
	IIterator[T any] interface {
		Value() T
		Valid() bool
		Next() IIterator[T]
	}

	IKVIterator[K, V any] interface {
		IIterator[V]
		Key() K
	}
)
