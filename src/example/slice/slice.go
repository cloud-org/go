package main

import "fmt"

func main()  {
	s := make([]int, 10)
	fmt.Println(len(s), cap(s))
	ss := s[10:]
	fmt.Println(len(ss), cap(s))
}