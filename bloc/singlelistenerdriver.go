package bloc

import "log"

// SingleListenerDriver is a bloc driver that supports a single subscriber.
// Calls to subscriber will return the same subscription.
type SingleListenerDriver struct {
	bloc      Bloc
	eventChan chan Event
	stateChan chan State
	subscription *singleListenerDriverSubscription
}

func NewSingleListenerDriver(bloc Bloc) *SingleListenerDriver {
	sld := &SingleListenerDriver{
		bloc:      bloc,
		eventChan: make(chan Event, 1),
		stateChan: make(chan State, 1),
	}
	sld.subscription = &singleListenerDriverSubscription{sld.stateChan, sld}
	sld.launch()
	return sld
}

func (sb *SingleListenerDriver) launch() {
	sb.stateChan <- sb.bloc.InitState()

	// Data provider
	go func() {
		if launchBloc, isLaunchBloc := sb.bloc.(BlocSelfLaunch); isLaunchBloc {
			launchBloc.OnLaunch(sb.stateChan)
		}

		for event := range sb.eventChan {
			log.Println("Reading event: ", event)
			sb.bloc.OnEvent(event, sb.stateChan)
		}
		close(sb.stateChan)
	}()
}

func (sb *SingleListenerDriver) Subscribe() Subscription {
	return sb.subscription
}

// PushEvent pushes an event to the bloc.  If the block is already handling an event, the event will get lost.
func (sb *SingleListenerDriver) PushEvent(event Event) {
	sb.eventChan <- event
}

func (sb *SingleListenerDriver) Close() {
	close(sb.eventChan)
}

type singleListenerDriverSubscription struct {
	c      chan State
	driver *SingleListenerDriver
}

func (s singleListenerDriverSubscription) Chan() chan State {
	return s.c
}

func (s singleListenerDriverSubscription) Close() {
}

