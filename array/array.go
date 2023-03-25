package array

import "fmt"

type (
	Array[T any] struct {
		elements []T
	}

	IArray[T any] interface {
		Assign(elem T)
		Fill(elem T)
		Size() int
		Capacity() int
		Empty() bool
		At(pos int) T
		Insert(pos int, elem T)
		Front() T
		Back() T
		Data() []T
		ToString() string
		Begin() *ArrayIterator[T]
		End() *ArrayIterator[T]
	}
)

func New[T any](size int) *Array[T] {
	return &Array[T]{elements: make([]T, size, size)}
}

func Clone[T any](other *Array[T]) *Array[T] {
	a := &Array[T]{elements: make([]T, other.Size(), other.Capacity())}
	for i := range other.elements {
		a.elements[i] = other.elements[i]
	}
	return a
}

func (a *Array[T]) Assign(elem T) {
	for i := range a.elements {
		a.elements[i] = elem
	}
}

func (a *Array[T]) Fill(elem T) {
	for i := range a.elements {
		a.elements[i] = elem
	}
}

func (a *Array[T]) Size() int {
	return len(a.elements)
}

func (a *Array[T]) Capacity() int {
	return cap(a.elements)
}

func (a *Array[T]) Empty() bool {
	return len(a.elements) == 0
}

func (a *Array[T]) At(pos int) T {
	if pos < 0 || pos >= len(a.elements) {
		panic("out of array range index")
	}
	return a.elements[pos]
}

func (a *Array[T]) Insert(pos int, elem T) {
	if pos < 0 || pos >= len(a.elements) {
		panic("out of array range index")
	}
	a.elements[pos] = elem
}

func (a *Array[T]) Front() T {
	return a.At(0)
}

func (a *Array[T]) Back() T {
	return a.At(len(a.elements) - 1)
}

func (a *Array[T]) Data() []T {
	return a.elements
}

func (a *Array[T]) ToString() string {
	return fmt.Sprintf("%v", a.elements)
}

func (a *Array[T]) Begin() *ArrayIterator[T] {
	return &ArrayIterator[T]{array: a, pos: 0}
}

func (a *Array[T]) End() *ArrayIterator[T] {
	return &ArrayIterator[T]{array: a, pos: a.Size()}
}
