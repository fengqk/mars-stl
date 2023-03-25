package llist

import (
	"github.com/fengqk/mars-stl/iterator"
)

type (
	ListIterator[T any] struct {
		node *Node[T]
	}

	IListIterator[T any] interface {
		iterator.IIterator[T]
		Prev() iterator.IIterator[T]
	}
)

func (iter *ListIterator[T]) Value() T {
	if iter.node == nil {
		panic("invalid iterator")
	}
	return iter.node.value
}

func (iter *ListIterator[T]) Valid() bool {
	return iter.node != nil
}

func (iter *ListIterator[T]) Next() iterator.IIterator[T] {
	if iter.node != nil {
		iter.node = iter.node.Next()
	}
	return iter
}

func (iter *ListIterator[T]) Prev() iterator.IIterator[T] {
	if iter.node != nil {
		iter.node = iter.node.Prev()
	}
	return iter
}
