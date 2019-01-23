package main

import (
	"github.com/dave/flux"
	"github.com/gopherjs/vecty"
	"github.com/kooksee/vvv/router"
)

func NewApp() *App {
	return &App{
	}
}

type App struct {
	vecty.Core

	dispatcher flux.DispatcherInterface
	watcher    flux.WatcherInterface
	notifier   flux.NotifierInterface
	stores     []flux.StoreInterface
	router     *router.Router
}

func (a *App) AddRoute(route string, c router.Handler) *App {
	a.router.HandleFunc(route, c)
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

func (a *App) Watch(key interface{}, f func(done chan struct{})) {
	a.watcher.Watch(key, f)
}

func (a *App) Delete(key interface{}) {
	a.watcher.Delete(key)
}

func (a *App) Stop() {
	a.router.Stop()
}

func (a *App) Navigate(path string) {
	a.router.Navigate(path)
}

func (a *App) Start() {
	a.router.Start()
}
