//package main
//
//import "fmt"
//
//func main() {
//
//	if 7%2 == 0 {
//		fmt.Println("7 is even")
//	} else {
//		fmt.Println("7 is odd")
//	}
//
//	if 8%4 == 0 {
//		fmt.Println("8 is divisible by 4")
//	}
//
//	if num := 9; num < 0 {
//		fmt.Println(num, "is negative")
//	} else if num < 10 {
//		fmt.Println(num, "has 1 digit")
//	} else {
//		fmt.Println(num, "has multiple digits")
//	}
//}
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println(111)
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} else {
		fmt.Println("1111")
	}
	if i := 9; i < 0 {
		fmt.Println("1111")
	} else if i < 10 {
		fmt.Println(i, "has 1 digit")
	} else {
		fmt.Println(1)
	}
}
