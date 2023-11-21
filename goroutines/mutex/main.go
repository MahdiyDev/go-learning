package main

import (
	"fmt"
	"sync"
	"time"
)

// import "fmt"

// func add(i *int) *int {
// 	*i += 12
// 	return i
// }

// func sum() *int {
// 	i := 1
// 	defer add(&i)
// 	return &i
// }

func inc(i *int, c *sync.Mutex) {
	c.Lock()
	fmt.Println("entered")
	for j := 0; j < 10; j++ {
		*i += 1
	}
	time.Sleep(time.Second)
	c.Unlock()
}

func main() {
	i := 12

	c := sync.Mutex{}

	go inc(&i, &c)
	go inc(&i, &c)
	go inc(&i, &c)
	go inc(&i, &c)

	time.Sleep(time.Second*5)

	fmt.Println(i)

	// fmt.Println(*sum())
}
