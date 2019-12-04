package formutils

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/event"
)

func BindStringValue(c vecty.Component, str *string) vecty.MarkupList {
	return vecty.Markup(
		vecty.Attribute("value", *str),
		event.Input(func(l *vecty.Event) {
			*str = l.Target.Get("value").String()
			vecty.Rerender(c)
		}),
	)
}
