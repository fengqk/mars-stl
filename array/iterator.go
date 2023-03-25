package array

import (
	"github.com/fengqk/mars-stl/iterator"
)

type (
	ArrayIterator[T any] struct {
		array *Array[T]
		pos   int
	}

	IArrayIterator[T any] interface {
		iterator.IIterator[T]
	}
)

func (iter *ArrayIterator[T]) Value() T {
	return iter.array.At(iter.pos)
}

func (iter *ArrayIterator[T]) Valid() bool {
	if iter.pos >= 0 && iter.pos < iter.array.Size() {
		return true
	}
	return false
}

func (iter *ArrayIterator[T]) Next() iterator.IIterator[T] {
	if iter.pos < iter.array.Size() {
		iter.pos++
	}
	return iter
}
