package llist

import (
	"errors"
	"fmt"
)

var ErrorOutOfRange = errors.New("Out of Range")

type (
	Node[T any] struct {
		value T
		list  *List[T]
		prev  *Node[T]
		next  *Node[T]
	}

	INode[T any] interface {
		Prev() *Node[T]
		Next() *Node[T]
	}

	List[T any] struct {
		head *Node[T]
		size int
	}

	IList[T any] interface {
		Size() int
		Empty() bool
		InsertBefore(elem T, at *Node[T]) *Node[T]
		InsertAfter(elem, at *Node[T]) *Node[T]
		PushBack(elem T) *Node[T]
		PushFront(elem T)
		Remove(elem *Node[T]) *Node[T]
		Clear()
		Front() T
		Back() T
		ToString() string
		Head() *ListIterator[T]
		Tail() *ListIterator[T]
	}
)

func New[T any]() *List[T] {
	return &List[T]{}
}

func (n *Node[T]) Prev() *Node[T] {
	if n.list == nil {
		return nil
	}
	if n == n.list.head {
		return nil
	}
	return n.prev
}

func (n *Node[T]) Next() *Node[T] {
	if n.list == nil {
		return nil
	}
	if n.next == n.list.head {
		return nil
	}
	return n.next
}

func (a *List[T]) Size() int {
	return a.size
}

func (a *List[T]) Empty() bool {
	return a.size == 0
}

func (a *List[T]) InsertBefore(elem T, at *Node[T]) *Node[T] {
	n := a.InsertAfter(&Node[T]{value: elem, list: a}, at.prev)
	if at == a.head {
		a.head = n
	}
	return n
}

func (a *List[T]) InsertAfter(elem, at *Node[T]) *Node[T] {
	elem.next = at.next
	elem.prev = at
	at.next.prev = elem
	at.next = elem
	a.size++
	return elem
}

func (a *List[T]) PushBack(elem T) *Node[T] {
	n := &Node[T]{value: elem, list: a}
	if a.size == 0 {
		n.prev = n
		n.next = n
		a.head = n
		a.size++
		return n
	}
	return a.InsertAfter(n, a.head.prev)
}

func (a *List[T]) PushFront(elem T) {
	n := a.PushBack(elem)
	a.head = n
}

func (a *List[T]) Remove(elem *Node[T]) *Node[T] {
	if elem == a.head {
		a.head = a.head.next
	}
	elem.prev.next = elem.next
	elem.next.prev = elem.prev
	elem.next = nil
	elem.prev = nil
	elem.list = nil
	a.size--
	if a.size == 0 {
		a.head = nil
	}
	return elem
}

func (a *List[T]) Clear() {
	a.head = nil
	a.size = 0
}

func (a *List[T]) Front() T {
	if a.size <= 0 {
		panic(ErrorOutOfRange)
	}
	return a.head.value
}

func (a *List[T]) Back() T {
	if a.size <= 0 {
		panic(ErrorOutOfRange)
	}
	return a.head.prev.value
}

func (a *List[T]) ToString() string {
	str := "["
	for iter := a.Head(); iter.Valid(); iter.Next() {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", iter.Value())
	}
	str += "]"
	return str
}

func (a *List[T]) Head() *ListIterator[T] {
	return &ListIterator[T]{node: a.head}
}

func (a *List[T]) Tail() *ListIterator[T] {
	if a.head == nil {
		return &ListIterator[T]{node: nil}
	}
	return &ListIterator[T]{node: a.head.prev}
}
