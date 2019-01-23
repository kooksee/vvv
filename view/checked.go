package view

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type _checked struct {
	_checked *bool
	fn       []func(bool)
}

func (t *_checked) Apply(h *vecty.HTML) {
	prop.Type(prop.TypeCheckbox).Apply(h)
	prop.Checked(*t._checked).Apply(h)
	event.Change(func(e *vecty.Event) {
		*t._checked = e.Target.Get("value").Bool()
		for _, _fn := range t.fn {
			_fn(*t._checked)
		}
	}).Apply(h)
}

func BindChecked(_check *bool, fn ...func(s bool)) *_checked {
	return &_checked{
		_checked: _check,
		fn:       fn,
	}
}
