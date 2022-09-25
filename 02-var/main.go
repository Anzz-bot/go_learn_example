package main

import (
	"fmt"
	"math"
)


func main() {
	var a string = "initial"
	var b int = 1
	var c = 2
	var d = true
	var e float64
	var f = float32(e)
	g := a + "foo"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)
	const s string = "constant"
	const h = 500000000
	j := 6e+11
	fmt.Println(s, h, j, math.Sin(h), math.Sin(j))

}
