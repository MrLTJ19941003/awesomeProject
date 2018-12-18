package main

import "fmt"

const (
	name  = "sss"
	i=0
	)

func read() (message string,error string){
	return "内容",""
}

func main() {
	var i1 int
	var i2 int8
	var i3 int16
	var i4 int32 = '字'
	var i5 int64
	fmt.Print(i1,i2,i3,i4,i5)
	fmt.Println("name：",name,"，i：",i)

	if name == "sss"{
		fmt.Println("成立")
	}else{
		fmt.Println("不成立")
	}

	if message ,error := read(); error != "" {
		fmt.Println(error)
	}else{
		fmt.Println(message)
	}
	switch {
		case name=="sss":
			fmt.Println("成立 s")
		default:
			fmt.Println("未知")
	}

	for i := 0;i<10;i++{
		fmt.Println("i : " ,i)
	}

}
