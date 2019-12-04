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
	OnEvent(event Event, transition chan State)
}

type BlocSelfLaunch interface {
	Bloc
	OnLaunch(transition chan State)
}

type Subscription interface {
	Chan() chan State
	Close()
}
