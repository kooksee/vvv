package components

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/style"
	"github.com/kooksee/vvv/example/todomvc/actions"
	"github.com/kooksee/vvv/example/todomvc/dispatcher"
	"github.com/kooksee/vvv/example/todomvc/store/model"
	"github.com/kooksee/vvv/view"
)

// ItemView is a vecty.Component which represents a single item in the TODO
// list.
type ItemView struct {
	vecty.Core

	Index     int         `vecty:"prop"`
	Item      *model.Item `vecty:"prop"`
	editing   bool
	EditTitle string
	input     *vecty.HTML
}

// Key implements the vecty.Keyer interface.
func (p *ItemView) Key() interface{} {
	return p.Index
}

func (p *ItemView) onDestroy(event *vecty.Event) {
	dispatcher.Dispatch(&actions.DestroyItem{
		Index: p.Index,
	})
}

//func (p *ItemView) onToggleCompleted(event *vecty.Event) {
func (p *ItemView) onToggleCompleted(_c bool) {
	dispatcher.Dispatch(&actions.SetCompleted{
		Index:     p.Index,
		Completed: _c,
	})
}

func (p *ItemView) onStartEdit(event *vecty.Event) {
	p.editing = true
	p.EditTitle = p.Item.Title
	vecty.Rerender(p)
	p.input.Node().Call("focus")
}

func (p *ItemView) onPrint(d string) {
	fmt.Println(d)
}

func (p *ItemView) onEditInput(event *vecty.Event) {
	p.EditTitle = event.Target.Get("value").String()
	vecty.Rerender(p)
}

func (p *ItemView) onStopEdit(event *vecty.Event) {
	p.editing = false
	vecty.Rerender(p)
	dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Index,
		Title: p.EditTitle,
	})
}

// Render implements the vecty.Component interface.
func (p *ItemView) Render() vecty.ComponentOrHTML {
	p.input = elem.Input(vecty.Markup(
		vecty.Class("edit"),
		//view.BindValue(p, "EditTitle", p.onPrint),
		view.BindValue(&p.EditTitle, p.onPrint),
	))

	return elem.ListItem(
		vecty.Markup(
			vecty.ClassMap{
				"completed": p.Item.Completed,
				"editing":   p.editing,
			},
		),

		elem.Div(
			vecty.Markup(
				vecty.Class("view"),
			),

			elem.Input(vecty.Markup(
				vecty.Class("toggle"),
				view.BindChecked(&p.Item.Completed, p.onToggleCompleted),
				//prop.Type(prop.TypeCheckbox),
				//prop.Checked(p.Item.Completed),
				//event.Change(p.onToggleCompleted),
			)),

			elem.Label(
				vecty.Markup(
					event.DoubleClick(p.onStartEdit),
				),
				vecty.Text(p.Item.Title),
			),

			elem.Button(vecty.Markup(
				vecty.Class("destroy"),
				event.Click(p.onDestroy),
			)),
		),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onStopEdit).PreventDefault(),
			),
			p.input,
		),
	)
}
