package main

import (
	"fmt"
	"log"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"strings"
)
const name  = "f:/"//Users/13045/go/src/awesomeProject/
const filenameTest  = "svn"

type worker struct {
	in chan string
	done chan string
}

func searchByfilename(w worker){
	var pathName string
	for {
		pathName = <- w.in
		fmt.Println(pathName)
		go func(pathName string) {
			//fmt.Println("22222222222",pathName)
			dir_list, _ := ioutil.ReadDir(pathName)
			for _, v := range dir_list {
				if strings.Contains(v.Name(),filenameTest){
					fmt.Println("已找到",v.Name(),"文件，目录为：", pathName,)
					go func() {
						w.done <- v.Name()
						//close(w.done)
					}()
					//			break
				}
				if v.IsDir(){
					str := pathName+v.Name()+"/"
					w.in <- str
					//w.done <- true
				}
				//go func(name ,pathn string,isDir bool) {
				//	if isDir{
				//		str := pathn+name+"/"
				//		w.in <- str
				//		w.done <- true
				//	}
				//}(v.Name(),pathName,v.IsDir(),)
			}

		}(pathName)
	}
}

func createSearchByfilename()worker{
	work := worker{
		in : make(chan string),
		done : make(chan string),
	}
	go searchByfilename(work)
	return work

}

func initSearch(name string,works worker){

	dir_list, e := ioutil.ReadDir(name)
	//var channels [40]chan string
	//q := make(chan string)
	//q1 := make(chan string)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	//go searchByfilename(q,q1)

	//go searchByfilename(works)
	for _, v := range dir_list {
		//channels[i] = createSearchByfilename()
		//fmt.Println(i, "=", name,v.Name(),"是否是目录",v.IsDir())
		if v.IsDir(){
			str := name+v.Name()+"/"
			works.in <- str
		}
	}
	//<-works.done
}
//var inTE, outTE *walk.TextEdit
type MyMainWindow struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	textArea  *walk.TextEdit
	results   *walk.ListBox
}

func (mw *MyMainWindow) clicked() {
	model := []string{}
	/*word := mw.searchBox.Text()
	text := mw.textArea.Text()
	model := []string{}
	for _, i := range search(text, word) {
		model = append(model, fmt.Sprintf("%d检索成功", i))
	}*/
	//fmt.Println(mw.results.Model())
	works := createSearchByfilename()
	initSearch(name,works)
	go func() {
		for{
			v := <- works.done
			if v == ""{
				fmt.Printf("已关闭")
				break
			}else{
				fmt.Println("===========" ,<-works.done)
				model = append(model,v)
			}
			//fmt.Println("===========" ,<-works.done)
		}
		mw.results.SetModel(model)
		//mw.results.Model()
		//model = append(model, fmt.Sprintf("%s检索成功",<-works.done))
	}()
	//mw.results.SetModel(model)
	//mw.results.SetModel(model)

}

func main() {
	mw := &MyMainWindow{}
	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SearchBox",
		MinSize:  Size{400, 500},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					LineEdit{
						AssignTo: &mw.searchBox,
					},
					PushButton{
						Text:      "检索",
						OnClicked: mw.clicked,
					},
				},
			},
			TextEdit{
				AssignTo: &mw.textArea,
			},
			ListBox{
				AssignTo: &mw.results,
				Row:10,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

	/*MainWindow{
		Title:   "测试windows系统窗口",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "搜索",
				OnClicked: func() {
					initSearch(name)
				},
			},
		},
	}.Run()*/

	//t1 := time.Now()
	//initSearch(name)
	//
	//time.Sleep(time.Second*12)
	//t2 := time.Now()
	//fmt.Println("运行时间：",t2.Sub(t1))

}