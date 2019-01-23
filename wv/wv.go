package wv

import (
	"github.com/zserge/webview"
	"sync"
)

var one sync.Once
var vvv *wv

func get() *wv {
	if vvv.w == nil {
		panic("please init webview")
	}
	return vvv
}

func InitSettings(cfg webview.Settings) {
	cfg.ExternalInvokeCallback = rpcHandle
	vvv.w = webview.New(cfg)
}

func init() {
	one.Do(func() {
		vvv = &wv{
			_m: make(map[string]RPCHandle),
		}
	})
}

type wv struct {
	_m map[string]RPCHandle
	w  webview.WebView
}

func Exit() {
	get().w.Exit()
}

func Run() {
	get().w.Run()
}

func Dispatch(fn func(w webview.WebView)) {
	get().w.Dispatch(func() {
		fn(get().w)
	})
}

func Bind(name string, v interface{}) (sync func(), err error) {
	return get().w.Bind(name, v)
}

func InjectCSS(css string) {
	get().w.InjectCSS(css)
}

func Eval(js string) error {
	return get().w.Eval(js)
}
