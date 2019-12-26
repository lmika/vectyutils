package bloc

type State interface{}
type Event interface{}

type Driver interface {
	Subscribe() Subscription
	PushEvent(event Event)
	Close()
}

type Bloc interface {
	InitState() State
	Handle(transition chan<- State, event <-chan Event)
}

type Subscription interface {
	Chan() chan State
	Close()
}
