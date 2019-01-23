package wv

import "github.com/zserge/webview"

func MsgInfo(title string, arg string) string {
	return get().w.Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo, title, arg)
}

func MsgWarn(title string, arg string) string {
	return get().w.Dialog(webview.DialogTypeAlert, webview.DialogFlagWarning, title, arg)
}

func MsgError(title string, arg string) string {
	return get().w.Dialog(webview.DialogTypeAlert, webview.DialogFlagError, title, arg)
}
