package main

import (
	"fmt"
	"sync"
	"time"
)

// chan
// func main() {
// 	src := []int{1, 2, 3, 4, 5}
// 	dst := []int{}

// 	c := make(chan int)

// 	for _, s := range src {
// 		go func(s int, c chan int) {
// 			result := s * 2
// 			c <- result
// 		}(s, c)
// 	}

// 	for _ = range src {
// 		num := <-c
// 		dst = append(dst, num)
// 	}

// 	fmt.Println(dst)
// 	close(c)
// }

// sync.Mutex
func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	var mu sync.Mutex

	for _, s := range src {
		go func(s int) {
			result := s * 2
			mu.Lock()
			dst = append(dst, result)
			mu.Unlock()
		}(s)
	}

	time.Sleep(time.Second)
	fmt.Println(dst)
}
