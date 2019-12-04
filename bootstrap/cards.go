package bootstrap

import "github.com/gopherjs/vecty"

func Card(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("card", elems)
}

func CardHeader(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("card-header", elems)
}

func CardBody(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return divWithClass("card-body", elems)
}

func CardTitle(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return tagWithClass("h5", "card-title", elems)
}

func CardText(elems ...vecty.ComponentOrHTML) *vecty.HTML {
	return tagWithClass("p", "card-text", elems)
}