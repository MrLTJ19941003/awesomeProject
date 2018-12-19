package main

import (
	"awesomeProject/crawler/frontend/controller"
	"net/http"
)

func main() { //搜索web http server 启动入口
	//path := getCurrentDirectory()
	//fmt.Println(path)
	http.Handle("/", http.FileServer(http.Dir("C:/Users/13045/go/src/awesomeProject/crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("C:/Users/13045/go/src/awesomeProject/crawler/frontend/view/search.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

//func getCurrentDirectory() string {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return strings.Replace(dir, "\\", "/", -1)
//}
