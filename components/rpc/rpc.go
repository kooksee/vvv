package rest

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/kooksee/vvv/models"
	"reflect"
	"sync"
	"time"
)

var one sync.Once
var rpcmap map[string]interface{}
var rpcclearmap map[int]string

func init() {
	one.Do(func() {
		rpcclearmap = make(map[int]string)
		rpcmap = make(map[string]interface{})
		js.Global.Set("__rpcmap", rpcmap)
	})

	go func() {
		for {
			for k, v := range rpcclearmap {
				if int(time.Now().UnixNano()) > k {
					delete(rpcclearmap, k)
					delete(rpcmap, v)
				}
				time.Sleep(time.Second * 5)
			}
		}
	}()
}

func RpcCall(service string, data string, fn interface{}) {

	_t := reflect.TypeOf(fn)
	if _t.Kind() != reflect.Func {
		panic("not func type")
	}

	_fname := _t.String() + "_" + fmt.Sprintf("%d", time.Now().UnixNano())
	rpcmap[_fname] = fn
	rpcclearmap[int(time.Now().UnixNano())] = _fname

	js.Global.Get("external").Get("invoke").Invoke((&models.RpcCall{
		Service:  service,
		Data:     data,
		Callback: _fname,
	}).Encode())
}
