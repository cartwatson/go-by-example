package main

import (
	"fmt"
	"math"
)

const s string = "str"

func main() {
	fmt.Println(s)

	const a = 500000000
	fmt.Println(a)
	fmt.Printf("Type of a: %T\n", a)

	const d = 3e20 / a
	fmt.Printf("Type of d: %T\n", d)
	fmt.Println(d)
	fmt.Printf("Type of d: %T\n", d)

	fmt.Println(int64(d))
	fmt.Printf("Type of d: %T\n", d)
	fmt.Println(math.Sin(a))
}
