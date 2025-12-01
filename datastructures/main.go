package main

import (
	"fmt"
	// "slices"
)

func main() {

	fmt.Println("\n---ARRAYS--------------------")
	var a [5]int // size = 5
	fmt.Println("init:", a)

	a[4] = 100
	fmt.Println("   a:", a)
	fmt.Println("a[4]:", a[4])
	fmt.Println(" len:", len(a))

	// [...] lets compiler count array length
	// non-specified indexes are zeros
	b := [...]int{100, 3: 400, 500}
	fmt.Println(" dcl:", b)

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("  2d:", twoD)

	fmt.Println("\n---SLICES--------------------")
	var s []string

	for _, char := range "hello world" { // _ is position, it's parsing UTF-8
		s = append(s, string(char))
	}
	fmt.Println("   s:", s)

	l := s[0:5]
	fmt.Println(" sl1:", l)

	t := []string{"g", "h", "i"}
	fmt.Println(" dcl:", t)

	fmt.Println("\n---MAPS----------------------")

	//make(map[key-type]val-type)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(" map:", m)

	// access as normal, same as arrays etc
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m) // error here that clear doesn't exist, it does
	fmt.Println("map:", m)
}
