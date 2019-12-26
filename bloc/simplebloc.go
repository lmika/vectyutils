package bloc

func BlocFn(initState State, handler func(transition chan<- State, event <-chan Event)) Bloc {
	return simpleBloc{initState, handler}
}

type simpleBloc struct {
	initState State
	handler func(transition chan<- State, event <-chan Event)
}

func (sb simpleBloc) InitState() State {
	return sb.initState
}

func (sb simpleBloc) Handle(transition chan<- State, event <-chan Event) {
	sb.handler(transition, event)
}
