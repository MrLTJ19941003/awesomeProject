package main

import "fmt"

//map --> hashMap
/*
map的key （限制）
1、map使用哈希表，必须可以比较相等
2、除了slice、map、function 的内建类型都可以作为Key
3、struct 类型不包含上述类型，也可作为Key
 */
func main() {
	map_addordelor()

}
func map_addordelor(){
	m := map[string]string{
		"name":"liu",
		"age":"23",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m,m2,m3)
	fmt.Println("Traversing map")//遍历map 注：遍历是无序的
	for k,v := range m{
		fmt.Println("key , value = " ,k,v)
	}
	//
	fmt.Println("getting values")//获取值
	if ages,ok := m["age"] ;ok{
		fmt.Println("ages values is ",ages)
	}else{
		fmt.Println("this key does not exist!")
	}
	//删除值
	fmt.Println("Deleting values")
	name,ok := m["name"]
	fmt.Println(name,ok)
	delete(m,"name")
	name,ok = m["name"]
	fmt.Println(name,ok)
}
