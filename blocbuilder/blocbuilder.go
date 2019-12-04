package blocbuilder

import (
	"github.com/gopherjs/vecty"
	"github.com/lmika/vectyutils/bloc"
)

func New(driver bloc.Driver, builder func(state bloc.State) vecty.ComponentOrHTML) *BlocBuilder {
	return &BlocBuilder{
		Bloc: driver,
		Builder: builder,
	}
}

type BlocBuilder struct {
	vecty.Core
	Bloc          bloc.Driver                                  `vecty:"prop"`
	Builder       func(state bloc.State) vecty.ComponentOrHTML `vecty:"prop"`
	OnSubscribe   func(initState bloc.State)                   `vecty:"prop"`
	OnUnsubscribe func()                                       `vecty:"prop"`

	//transitionIndex int
	state     bloc.State
	stateChan bloc.Subscription
	comp      vecty.ComponentOrHTML
}

func (p *BlocBuilder) Render() vecty.ComponentOrHTML {
	if p.state == nil {
		p.blocState()
	}

	return p.Builder(p.state)
}

func (b *BlocBuilder) blocState() bloc.State {
	if b.stateChan == nil {
		b.stateChan = b.Bloc.Subscribe()
		b.state = <-b.stateChan.Chan()

		go func() {
			//runtime.LockOSThread()
			for state := range b.stateChan.Chan() {
				b.state = state
				vecty.Rerender(b)
			}
		}()

		if b.OnSubscribe != nil {
			b.OnSubscribe(b.state)
		}
	}

	return b.state
}

func (b *BlocBuilder) Unmount() {
	b.stateChan.Close()
	b.stateChan = nil

	if b.OnUnsubscribe != nil {
		b.OnUnsubscribe()
	}
}
