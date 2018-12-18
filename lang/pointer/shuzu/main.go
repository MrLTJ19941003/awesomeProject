package main

import "fmt"

func swap(a,b *int) {
	*a , *b = *b ,*a
}

func printArray( arr []int)  {//数组默认值传递
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}

}

//指针不能运算，指针分为值类型、指针类型
func main() {
	/*a ,b := 3,11
	swap(&a,&b)
	fmt.Println(a,b)*/
	var arr1 [5]int
	var arr2 = [3]int{1,2,3}
	var arr3 = [...]int{2,3,5,6,8}
	var arr4 [4][5]int

	fmt.Println(arr1,arr2,arr3,arr4)
	/*for i:= range arr3{
		fmt.Println(i)
	}*/
	fmt.Println("printArray(arr3)")
	printArray(arr3[:])
	fmt.Println(arr3)
}
