
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
