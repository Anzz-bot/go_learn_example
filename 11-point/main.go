//package main
//
//import "fmt"
//
//func add2(n int) {
//	n += 2
//}
//
//func add2ptr(n *int) {
//	*n += 2
//}
//
//func main() {
//	n := 5
//	add2(n)
//	fmt.Println(n) // 5
//	add2ptr(&n)
//	fmt.Println(n) // 7
//}
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
