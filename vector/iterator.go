package vector

import (
	"github.com/fengqk/mars-stl/iterator"
)

type (
	VectorIterator[T any] struct {
		vector *Vector[T]
		pos    int
	}

	IVectorIterator[T any] interface {
		iterator.IIterator[T]
	}
)

func (iter *VectorIterator[T]) Value() T {
	return iter.vector.At(iter.pos)
}

func (iter *VectorIterator[T]) Valid() bool {
	if iter.pos >= 0 && iter.pos < iter.vector.Size() {
		return true
	}
	return false
}

func (iter *VectorIterator[T]) Next() iterator.IIterator[T] {
	if iter.pos < iter.vector.Size() {
		iter.pos++
	}
	return iter
}
