package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int , op string) (int ,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q, _ := div(a, b)
		return q,nil;
	default:
		return 0,fmt.Errorf("unsupported operation" + op);

	}
}

func div(a,b int) (q , r int){
	return a/b,a%b
}

func apply(op func(int,int) int,a,b int) int{
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d , %d)",opName,a ,b)
	return op(a,b)
}

func Pow(a,b int)  int{
	return int(math.Pow(float64(a),float64(b)))

}

func sum (number ...int) int {
	sum := 0
	for i:= range number{
		sum += i
	}
	return sum
}

func main() {
	if result,err := eval(3,4,"*") ;err!=nil{
		fmt.Println("error : " ,err)
	}else{
		fmt.Println("result : " ,result)
	}

	q, r := div(13, 3)
	fmt.Println(q,r)

	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i),float64(i2)))
	},3,4))

	fmt.Println(sum(1,2,3,4,5,6))
}
