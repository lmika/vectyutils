package bootstrap

import (
	"github.com/gopherjs/vecty"
	"strconv"
)

func Row(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("row", elems)
}

func Col(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("col-sm", elems)
}

func ColEx(colOptions ColOptions, elems ...vecty.ComponentOrHTML) *vecty.HTML {
	if colOptions.Width > 0 {
		return divWithClass("col-" + strconv.Itoa(colOptions.Width), elems)
	}

	return Col(elems...)
}

type ColOptions struct {
	Width	int
}