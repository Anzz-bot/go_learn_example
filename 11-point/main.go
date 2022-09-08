
package main

import "fmt"

func add(a int) {
	a += 2
}
func addPtr(a *int) {
	*a += 2
}

func main() {
	a := 2
	add(a)
	fmt.Println(a)
	addPtr(&a)
	fmt.Println(a)
}
