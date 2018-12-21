// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Start(outlist chan string) {
	mw := &MyMainWindow{model: NewEnvModel(),outList:outlist}
	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "实用搜索工具 1.0.0003_alpha测试版",
		MenuItems: []MenuItem{
			Menu{
				Text: "&编辑",
				Items: []MenuItem{
					Separator{},
					Action{
						Text:        "退出",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
		},
		MinSize: Size{650, 400},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			Composite{

				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					Label{Text: "关键词: "},
					LineEdit{
						AssignTo: &mw.keywords,
						Text:     "golang",
					},
					PushButton{
						Text:      "search",
						OnClicked: mw.GetList,
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 1, Spacing: 10},
				Children: []Widget{
					ListBox{
						AssignTo:        &mw.lb,
						Model:           mw.model,
						OnItemActivated: mw.lb_ItemActivated,
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	model    *EnvModel
	keywords *walk.LineEdit
	lb       *walk.ListBox
	outList  chan string
}

func (mw *MyMainWindow) lb_CurrentIndexChanged() {
	i := mw.lb.CurrentIndex()
	item := &mw.model.items[i]
	fmt.Println("CurrentIndex: ", i)
	fmt.Println("CurrentEnvVarName: ", item.filename)
}

func (mw *MyMainWindow) lb_ItemActivated() {
	value := mw.model.items[mw.lb.CurrentIndex()].filename
	walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
}

type EnvItem struct {
	filename string
}

type EnvModel struct {
	walk.ListModelBase
	items []EnvItem
}

func NewEnvModel() *EnvModel {
	m := &EnvModel{items: make([]EnvItem, 0)}

	return m
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index].filename
}

func (mw *MyMainWindow) GetList() {
	go func() {
		for{
			select {
			case str := <- mw.outList:
				items := []EnvItem{}
				m := &EnvModel{items: items}
				m.items = append(m.items, EnvItem{str})
				mw.lb.SetModel(m)
				mw.model = m
			}
		}

	}()

}
