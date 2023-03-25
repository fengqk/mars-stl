package main

import (
	"fmt"

	"github.com/fengqk/mars-stl/array"
)

func main() {
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
