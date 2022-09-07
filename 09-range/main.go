
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
