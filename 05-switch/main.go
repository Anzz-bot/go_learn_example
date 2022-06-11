//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//
//	a := 2
//	switch a {
//	case 1:
//		fmt.Println("one")
//	case 2:
//		fmt.Println("two")
//	case 3:
//		fmt.Println("three")
//	case 4, 5:
//		fmt.Println("four or five")
//	default:
//		fmt.Println("other")
//	}
//
//	t := time.Now()
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("It's before noon")
//	default:
//		fmt.Println("It's after noon")
//	}
//}

package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	}
	var tm = time.Now()
	switch {
	case tm.Hour()-12 < 0:
		fmt.Println("shangwu")
	default:
		fmt.Println("xiawu")

	}
}
