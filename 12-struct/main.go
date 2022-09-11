
package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	var a = user{name: "hxl", password: "6626"}
	b := user{name: "hxl", password: "6626"}
	c := user{name: "hxl"}
	c.password = "1234"
	fmt.Println(a, b, c)
	fmt.Println(checkPasswd(a, "1234"))
	fmt.Println(checkPasswd1(&a, "2345"))
	fmt.Println(a.password)
}
func checkPasswd(a user, password string) bool {
	return a.password == password
}
func checkPasswd1(a *user, password string) bool {
	a.password = "abcd"
	return a.password == password
}
