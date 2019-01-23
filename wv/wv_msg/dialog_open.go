package wv_msg

import "github.com/zserge/webview"

func FileOpen(w webview.WebView, title string, arg string) string{
	return w.Dialog(webview.DialogTypeOpen, webview.DialogFlagFile, title, arg)
}

func FileDir(w webview.WebView, title string, arg string)string {
	return w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, title, arg)
}

func FileSave(w webview.WebView, title string, arg string) string{
	return w.Dialog(webview.DialogTypeSave, webview.DialogFlagFile, title, arg)
}
