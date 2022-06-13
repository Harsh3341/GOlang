package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() {
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	go func() {
// 		i := 42
// 		ch <- i
// 		i = 27
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }
// ######################################################

func main() {
	ch := make(chan int)
	for j := 0; j < 10; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
