package main

import (
	"fmt"

	"github.com/fengqk/mars-stl/array"
	"github.com/fengqk/mars-stl/llist"
	"github.com/fengqk/mars-stl/vector"
)

func main() {

	if false {
		a := array.New[int](10)
		a.Fill(0)
		for i := 0; i < a.Size(); i++ {
			a.Insert(i, i)
		}
		fmt.Println(a.ToString())

		for iter := a.Begin(); iter.Valid(); iter.Next() {
			fmt.Println(iter.Value())
		}
	}

	if false {
		v := vector.New[int]()
		for i := 0; i < 16; i++ {
			v.PushBack(i)
		}
		fmt.Println(v.ToString())

		for iter := v.Begin(); iter.Valid(); iter.Next() {
			fmt.Println(iter.Value())
		}
	}

	if true {
		l := llist.New[int]()
		l.PushBack(2)
		l.PushBack(3)
		l.PushBack(4)
		l.PushFront(1)
		l.PushFront(0)
		fmt.Println(l.ToString())

		for iter := l.Head(); iter.Valid(); iter.Next() {
			fmt.Println(iter.Value())
		}
		for iter := l.Tail(); iter.Valid(); iter.Prev() {
			fmt.Println(iter.Value())
		}
	}
}
