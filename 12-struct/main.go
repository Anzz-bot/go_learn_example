//package main
//
//import "fmt"
//
//type user struct {
//	name     string
//	password string
//}
//
//func main() {
//	a := user{name: "wang", password: "1024"}
//	b := user{"wang", "1024"}
//	c := user{name: "wang"}
//	c.password = "1024"
//	var d user
//	d.name = "wang"
//	d.password = "1024"
//
//	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
//	fmt.Println(checkPassword(a, "haha"))   // false
//	fmt.Println(checkPassword2(&a, "haha")) // false
//}
//
//func checkPassword(u user, password string) bool {
//	return u.password == password
//}
//
//func checkPassword2(u *user, password string) bool {
//	return u.password == password
//}
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
