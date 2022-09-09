
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["unknow"])
	r, ok := m["unknow"]
	fmt.Println(r, ok)
	delete(m, "one")
	fmt.Println(m)

}
