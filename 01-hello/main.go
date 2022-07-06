package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}



////////////////func print(i int) {
////////////////	fmt.Println("hello world", i)
////////////////}
////////////////func main() {
////////////////	for i := 0; i < 5; i++ {
////////////////		go func(j int) {
////////////////			print(j)
////////////////		}(i)
////////////////	}
////////////////	time.Sleep(time.Second)
////////////////}
//////////////package main
//////////////
//////////////import "fmt"
//////////////
//////////////func main() {
//////////////	src := make(chan int)
//////////////	dest := make(chan int)
//////////////	go func() {
//////////////		defer close(src)
//////////////		for i := 0; i < 10; i++ {
//////////////			src <- i
//////////////		}
//////////////	}()
//////////////	go func() {
//////////////		defer close(dest)
//////////////		for i := range src {
//////////////			dest <- i
//////////////		}
//////////////	}()
//////////////	for i := range dest {
//////////////		fmt.Println(i)
//////////////	}
//////////////}
////////////package main
////////////
////////////import (
////////////	"fmt"
////////////	"sync"
////////////	"time"
////////////)
////////////
////////////var (
////////////	x    int64
////////////	lock sync.Mutex
////////////)
////////////
////////////func main() {
////////////	x = 0
////////////	for i := 0; i < 5; i++ {
////////////		go func() {
////////////			addWithLock()
////////////		}()
////////////	}
////////////	time.Sleep(time.Second)
////////////	fmt.Println(x)
////////////	x = 0
////////////	for i := 0; i < 5; i++ {
////////////		go func() {
////////////			addWithoutLock()
////////////		}()
////////////	}
////////////	time.Sleep(time.Second)
////////////	fmt.Println(x)
////////////}
////////////func addWithLock() {
////////////
////////////	for i := 0; i < 10000; i++ {
////////////		lock.Lock()
////////////		x += 1
////////////		lock.Unlock()
////////////	}
////////////}
////////////func addWithoutLock() {
////////////	for i := 0; i < 10000; i++ {
////////////		x += 1
////////////	}
////////////}
//////////
//////////package main
//////////
//////////import (
//////////	"fmt"
//////////	"sync"
//////////)
//////////
//////////func manygowait() {
//////////	var wg sync.WaitGroup
//////////	wg.Add(5)
//////////	for i := 0; i < 5; i++ {
//////////		go func(j int) {
//////////			defer wg.Done()
//////////			fmt.Println("hello", j)
//////////		}(i)
//////////	}
//////////	wg.Wait()
//////////}
//////////func main() {
//////////	manygowait()
//////////}
//////package main
//////
//////import "fmt"
//////
//////func pingfang() {
//////	src := make(chan int)
//////	tmp := make(chan int)
//////	go func() {
//////		defer close(src)
//////		for i := 0; i < 10; i++ {
//////			src <- i
//////		}
//////	}()
//////	go func() {
//////		defer close(tmp)
//////		for i := range src {
//////			tmp <- i * i
//////		}
//////	}()
//////	for i := range tmp {
//////		fmt.Println(i)
//////	}
//////}
//////func main() {
//////	src := make(chan int)
//////	i, ok := <-src
//////	fmt.Println(i, ok)
//////
//////}
//////
////package main
////
////import (
////	"fmt"
////	"sync"
////)
////
////var (
////	x    int64
////	lock sync.Mutex
////	wg   sync.WaitGroup
////)
////
////func addwithsync() {
////
////	for i := 0; i < 2000; i++ {
////		lock.Lock()
////		x++
////		lock.Unlock()
////	}
////}
////func addwithoutsync() {
////
////	for i := 0; i < 2000; i++ {
////		x++
////	}
////}
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
