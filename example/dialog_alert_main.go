package main

import (
	"fmt"
	"github.com/kooksee/vvv/wv"
	"github.com/zserge/webview"
)

type h1 struct {
	Name string
}

func (t *h1) Add() {
	t.Name = "hello"

}

func Hello(data string) (ret string, err error) {
	return data, nil
}

func main() {

	const (
		windowWidth  = 680
		windowHeight = 420
	)

	var indexHTML = `<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<style>* { margin: 0; padding: 0; box-sizing: border-box; }</style>
	</head>
	<body>
		<button onclick="test()">ddd</button>
		<script type="text/javascript">
			var __rpcmap={};
			__rpcmap["test1"]=function(d) {
					console.log(d);
			};
			__rpcmap["test2"]=function(d) {
					console.log("ssss");
			};

			function test() {
				console.log("hello");
				rpcCall("hello","test","test2");

				rpcCall("hello","test","test1");
			}
		</script>
	</body>
</html>
`

	wv.InitSettings(webview.Settings{
		Title:     "Loaded: Injected via JavaScript",
		Debug:     true,
		Resizable: true,
		Width:     windowWidth,
		Height:    windowHeight,
		//URL:       "data:text/html," + url.PathEscape(indexHTML),
		//URL: "http://localhost:8080/github.com/kooksee/vvv/example/todomvc/",
		URL: "http://localhost:8080/github.com/kooksee/vvv/example/html2vecty/",
	})

	fmt.Println(indexHTML)
	defer wv.Exit()
	wv.RPCRegistry("hello", Hello)

	//	wv.Dispatch(func(w webview.WebView) {
	//		a := wv.MsgInfo("Humble/examples", `
	//TodoMVC
	//TodoMVC is a Humble + GopherJS implementation of the famous TodoMVC application. It is finished and fully functional, implementing the entire TodoMVC spec. To learn more about the TodoMVC example, visit the README for TodoMVC.`)
	//		//a := wv_msg.MsgError(w, "he", "dd")
	//		//a := wv_msg.FileDir(w, "he", "dd")
	//		//a := wv_msg.FileOpen(w, "he", "dd")
	//		//a := wv_msg.FileSave(w, "he", "dd")
	//		fmt.Println(a)
	//	})
	wv.Run()
}
