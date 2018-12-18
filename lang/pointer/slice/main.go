package main

import "fmt"

func updateSlices(s []int){
	s[0] =100
}



//slice

//切片
func main() {
	/*arr:=[...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] = " ,arr[2:6])
	fmt.Println("arr[:6] = " ,arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = " ,s1)
	s2 := arr[:]
	fmt.Println("s2 = " ,s2)

	fmt.Println("After updateSlices s1")
	updateSlices(s1)
	fmt.Println("s1 = " ,s1)
	fmt.Println("arr = " ,arr)


	fmt.Println("After updateSlices s2")
	updateSlices(s2)
	fmt.Println("s2 = " ,s2)
	fmt.Println("arr = " ,arr)

	fmt.Println("reslice")
	fmt.Println("reslice 1 = " ,s2)
	s2 = s2[:5]
	fmt.Println("reslice 2 = " ,s2)
	s2 = s2[2:]
	fmt.Println("reslice 3 = " ,s2)
	//Extending slice
	fmt.Println("Extending slice")
	arr[0],arr[2] = 0,2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1 = %v ,len(s1) = %d,cap(s1) = %d ",s1,len(s1),cap(s1))
	fmt.Println()
	fmt.Printf("s2 = %v ,len(s2) = %d,cap(s2) = %d ",s2,len(s2),cap(s2))
	//slice 添加
	fmt.Println()
	fmt.Println("Append slice")
	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Printf("s3 = %d ,s4 = %d,s5 = %d ",s3,s4,s5)
	fmt.Println()
	fmt.Println(arr)
	fmt.Println("cap(arr) , cap(s4),cap(s5) = ",cap(arr),cap(s4),cap(s5))*/
	//createing slice创建切片
	var s []int
	for i:= 0;i<100 ;i++{
		printSlices(s)
		s = append(s, i*2 +1)
	}
	fmt.Println( "s = " , s)
	//s1 := []int{2,4,6,8}
	s1 := []int{2,4,6,8}
	printSlices(s1)
	s2 := make([]int,16)
	s3 := make([]int,10,32)
	printSlices(s2)
	printSlices(s3)
	//切片元素复制
	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlices(s2)

	//删除切片中的元素
	fmt.Println("Deleting elements from size")
	s2 = append(s2[:3],s2[4:]...)
	printSlices(s2)

	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front = ",front)
	printSlices(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail = ",tail)
	printSlices(s2)



}
func printSlices(s []int){
	fmt.Printf("%v ,len(s) = %d,cap(s) = %d" ,s ,len(s),cap(s))
	fmt.Println()
}
