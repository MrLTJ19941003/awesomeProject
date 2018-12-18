package main

import "fmt"

func addr() func(int) int  {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main(){
	a := addr()
	for i := 0;i<=10;i++{
		fmt.Println(a(i))
	}
}
