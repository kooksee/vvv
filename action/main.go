package main

import (
	"fmt"
	"github.com/dave/flux"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/kooksee/vvv/router"
	"strings"
)

func main() {
	app := NewApp()
	// 添加存储状态
	app.AddStore(NewEditorStore())

	// 添加路由
	app.AddRoute("/", func(ctx *router.Context) {
		vecty.RenderBody(&MainView{store: NewEditorStore()})
	})
	app.Start()
}

func NewEditorStore() *EditorStore {
	s := &EditorStore{
		html: strings.TrimSpace(""),
	}
	return s
}

type EditorStore struct {
	html, code string
}

func (s *EditorStore) Html() string {
	return s.html
}

func (s *EditorStore) Code() string {
	return s.code
}

func (s *EditorStore) Handle(payload *flux.Payload) bool {
	switch a := payload.Action.(type) {
	default:
		fmt.Println(a)
	}
	return true
}

// MainView is the parent level view.
type MainView struct {
	vecty.Core

	store *EditorStore
}

// Render returns a <body> element with the entire app inside of it.
func (pv *MainView) Render() vecty.ComponentOrHTML {
	return elem.Heading1(
		vecty.Markup(vecty.Class("blog-header")),
		vecty.Text("Blog-ish"),
	)
}
