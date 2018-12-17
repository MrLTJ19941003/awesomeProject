package main

import "fmt"

func main(){
	fmt.Println(MaxlengthString("abcabcbb"))
	fmt.Println(MaxlengthString("bbbb"))
	fmt.Println(MaxlengthString("pwwkew"))
	fmt.Println(MaxlengthString(""))
	fmt.Println(MaxlengthString("b"))
	fmt.Println(MaxlengthString("abcdefg"))
	fmt.Println(MaxlengthString("yes我喜欢慕课网"))

}


func MaxlengthString(s string) int {
	maxlength := 0
	lastOccurred := make(map[rune]int)
	start := 0

	for i,ch :=range []rune(s){
		lastI,ok := lastOccurred[ch]
		if ok && lastI >= start{
			start = lastOccurred[ch] + 1
		}
		if i-start+1 >maxlength{
			maxlength = i-start+1
		}
		lastOccurred[ch] = i
	}
	return maxlength

}
