//package main
//
//import "fmt"
//
//func main() {
//
//	s := make([]string, 3)
//	s[0] = "a"
//	s[1] = "b"
//	s[2] = "c"
//	fmt.Println("get:", s[2])   // c
//	fmt.Println("len:", len(s)) // 3
//
//	s = append(s, "d")
//	s = append(s, "e", "f")
//	fmt.Println(s) // [a b c d e f]
//
//	c := make([]string, len(s))
//	copy(c, s)
//	fmt.Println(c) // [a b c d e f]
//
//	fmt.Println(s[2:5]) // [c d e]
//	fmt.Println(s[:5])  // [a b c d e]
//	fmt.Println(s[2:])  // [c d e f]
//
//	good := []string{"g", "o", "o", "d"}
//	fmt.Println(good) // [g o o d]
//}
package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(len(s), cap(s))
	s = append(s, "d")
	fmt.Println(len(s), cap(s))
	s = append(s, "e", "f", "g")
	fmt.Println(len(s), cap(s))
	c := make([]string, len(s))
	m := s[:]
	copy(c, s)
	fmt.Println(m)
	s[1] = "t"
	fmt.Println(s, m)
	h := []int{}
	fmt.Println(h)

}
