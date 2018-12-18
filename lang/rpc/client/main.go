package main

import (
	"awesomeProject/lang/rpc"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn,err := net.Dial("tcp",":8999")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil {
		log.Printf("client error : %v " ,err)
	}
	fmt.Println(result)
}
