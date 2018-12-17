package main

import (
	"fmt"
	"time"
)

func main() {
	//a := []int{}
	for i:= 0 ;i<10;i++{
		go func(i int) {
			for{
				fmt.Printf("goroutine from %d \n" ,i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
