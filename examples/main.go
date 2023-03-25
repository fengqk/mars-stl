package main

import (
	"fmt"

	"github.com/fengqk/mars-stl/array"
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

	if true {
		v := vector.New[int]()
		for i := 0; i < 16; i++ {
			v.PushBack(i)
		}
		fmt.Println(v.ToString())

		for iter := v.Begin(); iter.Valid(); iter.Next() {
			fmt.Println(iter.Value())
		}
	}
}
