package router

import (
	"github.com/gopherjs/vecty"
	"github.com/lmika/vectyutils/bloc"
	"github.com/lmika/vectyutils/blocbuilder"
)

type Router struct {
	vecty.Core
	routeMatcher *simpleRouteMatcher
	driver       bloc.Driver
}

func New(configs ...Config) *Router {
	r := &Router{
		routeMatcher: new(simpleRouteMatcher),
	}

	for _, config := range configs {
		config(r)
	}

	r.routeMatcher.init()
	r.driver = bloc.NewSingleListenerDriver(&routerBloc{driver: defaultDriver})
	return r
}

func (r *Router) Key() interface{} {
	return r.routeMatcher
}

func (r *Router) Render() vecty.ComponentOrHTML {
	return blocbuilder.New(r.driver, func(state bloc.State) vecty.ComponentOrHTML {
		// TEMP
		path := r.interpretPath(state)

		pathMatch, context := r.routeMatcher.match(path)
		if pathMatch == nil {
			return nil
		}

		return pathMatch.builder(context)
	})
}

func (r *Router) interpretPath(state bloc.State) string {
	var path = ""
	if state != nil {
		path = state.(routeState).path
	}

	// If path is '' then default to "/"
	//if path == "" {
	//	path = "/"
	//} else if !strings.HasPrefix(path, "/") {
	//	path = "/" + path
	//}
	return path
}

type Context struct {
	components map[string]string
}

// Var returns the name of the path variable.  (or index)
func (rc *Context) Var(name string) string {
	return rc.components[name]
}

type builder func(ctx *Context) vecty.ComponentOrHTML

type Config func(r *Router)

func Route(pattern string, builder func(ctx *Context) vecty.ComponentOrHTML) func(r *Router) {
	return func(r *Router) {
		r.routeMatcher.addPath(pattern, builder)
	}
}
