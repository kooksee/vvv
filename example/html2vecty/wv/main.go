package main

import (
	"github.com/kooksee/vvv/wv"
	"github.com/zserge/webview"
)

type h struct {
	Name string
}

func (t *h) Add() {
	t.Name = "hh"
}

func main() {
	wv.InitSettings(webview.Settings{
		Title:     "Loaded: Injected via JavaScript",
		Debug:     true,
		Resizable: true,
		URL:       "http://localhost:8080/github.com/kooksee/vvv/example/html2vecty/static",
	})
	defer wv.Exit()
	wv.Run()
}
