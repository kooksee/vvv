package main

import (
	"github.com/dave/flux"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/kooksee/vvv/internal/utils"
	"marwan.io/vecty-router"
)

func NewApp() *App {
	return &App{
		_routers: make(map[string]int),
	}
}

type App struct {
	vecty.Core

	dispatcher flux.DispatcherInterface
	watcher    flux.WatcherInterface
	notifier   flux.NotifierInterface
	stores     []flux.StoreInterface
	routers    []vecty.MarkupOrChild
	_routers   map[string]int
}

func (a *App) NotFoundHandler(c vecty.Component) *App {
	utils.Assert(a._routers["__not_found"] == 1, "not found router had existed")
	a._routers["__not_found"] += 1
	a.routers = append(a.routers, router.NewRoute("__not_found", c, router.NewRouteOpts{ExactMatch: true}))
	return a
}

func (a *App) AddRoute(route string, c vecty.Component) *App {
	utils.Assert(a._routers[route] == 1, "%s had existed", route)

	a._routers[route] += 1
	a.routers = append(a.routers, router.NewRoute(route, c, router.NewRouteOpts{ExactMatch: true}))
	return a
}

func (a *App) AddStore(store flux.StoreInterface) {
	a.stores = append(a.stores, store)
}

func (a *App) Init() {

	n := flux.NewNotifier()
	a.notifier = n
	a.watcher = n

	a.dispatcher = flux.NewDispatcher(
		// Notifier:
		a.notifier,
		// Stores:
		a.stores...
	)
}

func (a *App) Dispatch(action flux.ActionInterface) chan struct{} {
	return a.dispatcher.Dispatch(action)
}
func (a *App) Render() vecty.ComponentOrHTML {
	return elem.Body(a.routers...)
}

func (a *App) Watch(key interface{}, f func(done chan struct{})) {
	a.watcher.Watch(key, f)
}

func (a *App) Delete(key interface{}) {
	a.watcher.Delete(key)
}

func (a *App) Run(title string) {
	vecty.SetTitle(title)
	vecty.RenderBody(a)
}
