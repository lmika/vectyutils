package bootstrap

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func MenuItem(text string, onClick func()) vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("dropdown-item"),
			vecty.Attribute("type", "button"),
			event.Click(func(i *vecty.Event) {
				if onClick != nil {
					onClick()
				}
			}),
		),
		vecty.Text(text),
	)
}
