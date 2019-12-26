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
	State     bloc.State			`vecty:"prop"`
	StateChan bloc.Subscription		`vecty:"prop"`
}

func (p *BlocBuilder) Render() vecty.ComponentOrHTML {
	if p.State == nil {
		p.blocState()
	}

	return p.Builder(p.State)
}

func (b *BlocBuilder) blocState() bloc.State {
	if b.StateChan == nil {
		b.StateChan = b.Bloc.Subscribe()
		b.State = <-b.StateChan.Chan()

		go func() {
			//runtime.LockOSThread()
			for state := range b.StateChan.Chan() {
				b.State = state
				vecty.Rerender(b)
			}
		}()

		if b.OnSubscribe != nil {
			b.OnSubscribe(b.State)
		}
	}

	return b.State
}

func (b *BlocBuilder) Unmount() {
	b.StateChan.Close()
	b.StateChan = nil

	if b.OnUnsubscribe != nil {
		b.OnUnsubscribe()
	}
}
