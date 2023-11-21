// package main

// // func sum(c chan int, x int) {
// // 	for i := 1; i < 1_000_000_00; i++ {
// // 		x += i
// // 	}
// // 	c <- x
// // }

// func main() {
// 	// c := make(chan int)
// 	// go sum(c, 10)
// 	// go sum(c, 21)
// 	// go sum(c, 22)
// 	// res, res1, res2 := <-c, <-c, <-c

// 	// fmt.Println(res, res1, res2)

// 	// ch := make(chan int, 2_000_000)

// 	// for i := 0; i < 2_000_000; i++ {
// 	// 	ch <- i + rand.Int()
// 	// }

// 	// for i := 0; i < 2_000_000; i++ {
// 	// 	fmt.Println(<-ch)
// 	// }

// 	// close(ch)

// 	// time.Sleep(time.Second * 5)
// }

package main

import (
	"fmt"
)

func fibonacci(c, quit chan int, hello chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-hello:
			fmt.Println("hello")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	hello := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		hello <- 123
		quit <- 0
	}()
	fibonacci(c, quit, hello)
	// time.Sleep(time.Second * 2)
}
