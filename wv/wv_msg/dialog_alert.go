package wv_msg

import "github.com/zserge/webview"

func MsgInfo(w webview.WebView, title string, arg string) string{
	return w.Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo, title, arg)
}

func MsgWarn(w webview.WebView, title string, arg string) string{
	return w.Dialog(webview.DialogTypeAlert, webview.DialogFlagWarning, title, arg)
}

func MsgError(w webview.WebView, title string, arg string) string{
	return w.Dialog(webview.DialogTypeAlert, webview.DialogFlagError, title, arg)
}
