package bootstrap

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Container(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("container", elems)
}

func FormGroup(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("form-group", elems)
}

func DivWithClass(cls string, elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return elem.Div(vecty.Markup(vecty.Class(cls)), vecty.List(elems))
}

func divWithClass(cls string, elems []vecty.ComponentOrHTML) *vecty.HTML {
	return elem.Div(vecty.Markup(vecty.Class(cls)), vecty.List(elems))
}

func divWithClasses(cls []string, elems []vecty.ComponentOrHTML) *vecty.HTML {
	return elem.Div(vecty.Markup(vecty.Class(cls...)), vecty.List(elems))
}

func tagWithClass(el, cls string, elems []vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag(el, vecty.Markup(vecty.Class(cls)), vecty.List(elems))
}