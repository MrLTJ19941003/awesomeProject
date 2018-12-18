package main

import (
	"awesomeProject/lang/retieveriever/mock"
	real2 "awesomeProject/lang/retieveriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,params map[string]string)string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url  = "http://www.imooc.com"
func seesion(retrieverPoster RetrieverPoster) string {
	retrieverPoster.Post(url, map[string]string{
		"content":"auther fask imooc",
	})
	return retrieverPoster.Get(url)
}

func download(r Retriever) string{
	return r.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"imocc"}
	r = &retriever
	inspect(r)
	r = &real2.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Millisecond,
	}
	inspect(r)


	seesion(&retriever)
	inspect(&retriever)
	//fmt.Println(download(r))
}

func inspect(r Retriever){
	fmt.Printf("%T %v\n",r,r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content : " ,v.Content)
	case *real2.Retriever:
		fmt.Println("UserAgent : " ,v.UserAgent)

	}
}
