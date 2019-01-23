package view

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type _input struct {
	name *string
	fn   []func(string)
}

func (t *_input) Apply(h *vecty.HTML) {
	prop.Value(*t.name).Apply(h)
	prop.Autofocus(true).Apply(h)
	event.Input(func(i *vecty.Event) {
		*t.name = i.Target.Get("value").String()
		for _, _fn := range t.fn {
			_fn(*t.name)
		}
	}).Apply(h)
}

func BindValue(name *string, fn ...func(s string)) *_input {
	return &_input{
		name: name,
		fn:   fn,
	}
}
