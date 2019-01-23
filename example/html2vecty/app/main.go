package main

import (
	"github.com/gopherjs/vecty"
)

func main() {
	app := &App{}
	app.Init()
	vecty.SetTitle("test")
	vecty.RenderBody(NewPage(app))
}
