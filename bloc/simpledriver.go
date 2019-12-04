package bloc

import "log"

type SimpleDriver struct {
	bloc          Bloc
	currState     State
	eventChan     chan Event
	subscribers   *subscriptionList
	stateListener chan State
}

func NewSimpleDriver(bloc Bloc) *SimpleDriver {
	sb := &SimpleDriver{
		bloc:        bloc,
		currState:   bloc.InitState(),
		eventChan:   make(chan Event),
		subscribers: newSubscriptionList(),
	}
	sb.launch()
	return sb
}

func (sb *SimpleDriver) launch() {
	sb.stateListener = make(chan State, 1)

	// Data provider
	go func() {
		if launchBloc, isLaunchBloc := sb.bloc.(BlocSelfLaunch); isLaunchBloc {
			launchBloc.OnLaunch(sb.stateListener)
		}

		for event := range sb.eventChan {
			log.Println("Reading event: ", event)
			sb.bloc.OnEvent(event, sb.stateListener)
		}
	}()

	// Fanout
	go func() {
		// XXX - modifying the listener channels causes a data race
		for s := range sb.stateListener {
			sb.currState = s
			sb.subscribers.notify(s)
		}
	}()
}

func (sb *SimpleDriver) Subscribe() Subscription {
	// DATA RACE HERE
	return sb.subscribers.add(sb.currState)
}

// PushEvent pushes an event to the bloc.  If the block is already handling an event, the event will get lost.
func (sb *SimpleDriver) PushEvent(event Event) {
	select {
	case sb.eventChan <- event:
	default:
		// Ignore (?)
	}
}

func (sb *SimpleDriver) Close() {
	sb.subscribers.close()
}
