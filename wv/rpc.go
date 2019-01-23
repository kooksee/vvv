package wv

import (
	"fmt"
	"github.com/kooksee/vvv/models"
	"github.com/kooksee/vvv/wv/wv_msg"
	"github.com/zserge/webview"
)

func RPCRegistry(name string, fn RPCHandle) {
	get()._m[name] = fn
}

func rpcHandle(w webview.WebView, data string) {
	t := get()

	fmt.Println(data)

	rpc := &models.RpcCall{}
	if err := rpc.Decode(data); err != nil {
		wv_msg.MsgError(t.w, "rpc call error", err.Error())
		return
	}

	if t._m[rpc.Service] == nil {
		wv_msg.MsgError(t.w, "rpc call func is nil", fmt.Sprintf("service %s not found", rpc.Service))
		return
	}

	dt, err := t._m[rpc.Service](rpc.Data)
	if err != nil {
		wv_msg.MsgError(t.w, "rpc callback exec error", err.Error())
		return
	}

	if err := t.w.Eval(fmt.Sprintf(`__rpcmap["%s"]("%s")`, rpc.Callback, dt)); err != nil {
		wv_msg.MsgError(t.w, "rpc callback eval error", err.Error())
		return
	}
}
