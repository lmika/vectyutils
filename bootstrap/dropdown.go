package bootstrap

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Dropdown struct {
	vecty.Core

	Variant		string			`vecty:"prop"`
	Text		string			`vecty:"prop"`
	SubItems	vecty.List		`vecty:"prop"`
}


func (d *Dropdown) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("btn-group"),
			vecty.Attribute("role", "group"),
		),
		elem.Button(
			vecty.Markup(
				vecty.Class("btn", d.Variant, "dropdown-toggle"),
				vecty.Attribute("type", "button"),
				vecty.Data("toggle", "dropdown"),
			),
			vecty.Text(d.Text),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("dropdown-menu"),
			),
			d.SubItems,
		),
	)
}
