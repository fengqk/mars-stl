package vector

import (
	"errors"
	"fmt"
)

var ErrorOutOfRange = errors.New("Out of Range")

type (
	Vector[T any] struct {
		elements []T
	}

	IVector[T any] interface {
		Size() int
		Capacity() int
		Empty() bool
		At(pos int) T
		Insert(pos int, elem T)
		PushBack(elem T)
		PopBack()
		Erase(pos int)
		EraseRange(spos, epos int)
		Reserve(capsize int)
		Clear()
		Front() T
		Back() T
		Data() []T
		ToString() string
		Begin() *VectorIterator[T]
		End() *VectorIterator[T]
	}
)

func New[T any]() *Vector[T] {
	return &Vector[T]{elements: make([]T, 0, 8)}
}

func Clone[T any](other *Vector[T]) *Vector[T] {
	a := &Vector[T]{elements: make([]T, other.Size(), other.Capacity())}
	for i := range other.elements {
		a.elements[i] = other.elements[i]
	}
	return a
}

func (a *Vector[T]) Size() int {
	return len(a.elements)
}

func (a *Vector[T]) Capacity() int {
	return cap(a.elements)
}

func (a *Vector[T]) Empty() bool {
	return len(a.elements) == 0
}

func (a *Vector[T]) At(pos int) T {
	if pos < 0 || pos >= len(a.elements) {
		panic(ErrorOutOfRange)
	}
	return a.elements[pos]
}

func (a *Vector[T]) Insert(pos int, elem T) {
	if pos < 0 || pos >= len(a.elements) {
		panic(ErrorOutOfRange)
	}
	a.elements = append(a.elements, elem)
	for i := len(a.elements) - 1; i > pos; i-- {
		a.elements[i] = a.elements[i-1]
	}
	a.elements[pos] = elem
}

func (a *Vector[T]) PushBack(elem T) {
	a.elements = append(a.elements, elem)
}

func (a *Vector[T]) PopBack() T {
	if a.Empty() {
		panic(ErrorOutOfRange)
	}
	elem := a.Back()
	a.elements = a.elements[:len(a.elements)-1]
	return elem
}

func (a *Vector[T]) Erase(pos int) {
	a.EraseRange(pos, pos+1)
}

func (a *Vector[T]) EraseRange(spos, epos int) {
	if spos > epos || spos < 0 || epos > a.Size() {
		return
	}
	lelems := a.elements[:spos]
	relems := a.elements[epos:]
	a.elements = append(lelems, relems...)
}

func (a *Vector[T]) Reserve(capsize int) {
	if cap(a.elements) >= capsize {
		return
	}
	elems := make([]T, len(a.elements), capsize)
	for i := 0; i < len(a.elements); i++ {
		elems[i] = a.elements[i]
	}
	a.elements = elems
}

func (a *Vector[T]) Clear() {
	a.elements = a.elements[:0]
}

func (a *Vector[T]) Front() T {
	return a.At(0)
}

func (a *Vector[T]) Back() T {
	return a.At(len(a.elements) - 1)
}

func (a *Vector[T]) Data() []T {
	return a.elements
}

func (a *Vector[T]) ToString() string {
	return fmt.Sprintf("%v", a.elements)
}

func (a *Vector[T]) Begin() *VectorIterator[T] {
	return &VectorIterator[T]{vector: a, pos: 0}
}

func (a *Vector[T]) End() *VectorIterator[T] {
	return &VectorIterator[T]{vector: a, pos: a.Size()}
}
