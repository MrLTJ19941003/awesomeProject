package main

import (
	"awesomeProject/lang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener,err := net.Listen("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	for {
		conn ,err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
