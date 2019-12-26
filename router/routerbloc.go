package router

import (
	"github.com/lmika/vectyutils/bloc"
)

type routeState struct {
	path string
}

type routeEvent struct {
	newPath string
}


type routerBloc struct{
	driver	*hashDriver
}

func (rb *routerBloc) InitState() bloc.State {
	return routeState{path: rb.driver.currentPath()}
}

func (rb *routerBloc) Handle(transition chan<- bloc.State, event <-chan bloc.Event) {
	pathChan := make(chan string)
	closeFn := rb.driver.subscribeToChanges(func() {
		pathChan <- rb.driver.currentPath()
	})
	defer closeFn()

	active := true
	for active {
		select {
		case _, ok := <-event:
			if !ok {
				active = false
			}
		case path := <-pathChan:
			transition <- routeState{path: path}
		}
	}
}