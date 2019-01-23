package wv_utils

import (
	"fmt"
	"github.com/kooksee/vvv/internal/utils"
	"github.com/zserge/webview"
)

func Back(w webview.WebView, ) {
	utils.AssertErr(w.Eval("window.history.back()"))
}

func Forward(w webview.WebView) {
	utils.AssertErr(w.Eval("window.history.forward()"))
}

func Reload(w webview.WebView) {
	utils.AssertErr(w.Eval(fmt.Sprintf("window.location.reload()")))
}

func Open(w webview.WebView, url string) {
	utils.AssertErr(w.Eval(fmt.Sprintf(`window.location.pathname = "%s"`, url)))
}
