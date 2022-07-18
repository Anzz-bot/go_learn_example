package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}


////func main() {
////	x = 0
////
////	wg.Add(5)
////	for i := 0; i < 5; i++ {
////		go func() {
////			addwithsync()
////			defer wg.Done()
////		}()
////	}
////	wg.Wait()
////	fmt.Println(x)
////	x = 0
////	wg.Add(5)
////	for i := 0; i < 5; i++ {
////		go func() {
////			addwithoutsync()
////			defer wg.Done()
////		}()
////
////	}
////	wg.Wait()
////	fmt.Println(x)
////
////}
//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a float32
//	var b float32
//	fmt.Println(reflect.TypeOf(a).Name())
//	fmt.Println(reflect.TypeOf(b).Name())
//	fmt.Println(a / b)
//}
