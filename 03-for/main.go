
//
//	for n := 0; n < 5; n++ {
//		if n%2 == 0 {
//			continue
//		}
//		fmt.Println(n)
//	}
//	for i <= 3 {
//		fmt.Println(i)
//		i = i + 1
//	}
//}
package main

import "fmt"

func main() {
	h := 1
	for {
		fmt.Println("loop")
		break
	}
	for i := 7; i < 9; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	for h <= 3 {
		fmt.Println(h)
		h++
	}
}
