package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("A: ", i)
		}()
	}

	ch := make(chan int)
	<-ch
	//time.Sleep(1 * time.Hour)
	//A:  9 runnext
	//A:  0
	//A:  1
	//A:  2
	//A:  3
	//A:  4
	//A:  5
	//A:  6
	//A:  7
	//A:  8
}
