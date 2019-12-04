package bloc

func SimpleBloc(initState State, eventHandler func(event Event, stateChan chan State)) Bloc {
	return simpleBloc{initState, eventHandler}
}

type simpleBloc struct {
	initState State
	eventHandler func(event Event, stateChan chan State)
}

func (sb simpleBloc) InitState() State {
	return sb.initState
}

func (sb simpleBloc) OnEvent(event Event, transition chan State) {
	sb.eventHandler(event, transition)
}
