//package main
//
//import "fmt"
//
//type user struct {
//	name     string
//	password string
//}
//
//func (u user) checkPassword(password string) bool {
//	return u.password == password
//}
//
//func (u *user) resetPassword(password string) {
//	u.password = password
//}
//
//func main() {
//	a := user{name: "wang", password: "1024"}
//	a.resetPassword("2048")
//	fmt.Println(a.checkPassword("2048")) // true
//}
package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u *user) resetPassword(password string) {
	u.password = password
}
func (u user) checkPassword(password string) bool {
	return u.password == password
}
func main() {
	a := user{
		name:     "hxl",
		password: "1234",
	}
	fmt.Println(a)
	a.checkPassword("11111")
	a.resetPassword("2456")
	fmt.Println(a.password)
}
