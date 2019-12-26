package router

import (
	"regexp"
	"sort"
	"strconv"
)

type simpleRouteMatcher struct {
	paths	[]*simpleRouteMatchPath
}

func (srm *simpleRouteMatcher) addPath(path string, builder builder) {
	p := newSimpleRouteMatchPath(path, builder)
	if p == nil {
		return
	}
	srm.paths = append(srm.paths, p)
}

func (srm *simpleRouteMatcher) init() {
	sort.Slice(srm.paths, func(i int, j int) bool {
		return len(srm.paths[i].origPath) > len(srm.paths[j].origPath)
	})
}

func (srm *simpleRouteMatcher) match(path string) (*simpleRouteMatchPath, *Context) {
	for _, p := range srm.paths {
		submatchers := p.regex.FindStringSubmatch(path)
		if submatchers != nil {
			rc := &Context{make(map[string]string)}
			for i, s := range submatchers[1:] {
				rc.components[strconv.Itoa(i)] = s
			}
			return p, rc
		}
	}
	return nil, nil
}


type simpleRouteMatchPath struct {
	origPath	string
	regex		*regexp.Regexp
	builder		builder
}

func newSimpleRouteMatchPath(origPath string, builder builder) *simpleRouteMatchPath {
	r, err := regexp.Compile(origPath)
	if err != nil {
		return nil
	}
	return &simpleRouteMatchPath{origPath, r, builder}
}