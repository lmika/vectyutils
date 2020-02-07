package blocbuilder

import (
	"github.com/gopherjs/vecty"
	"github.com/lmika/vectyutils/bloc"
)

func New(driver bloc.Driver, builder func(state bloc.State) vecty.ComponentOrHTML) *BlocBuilder {
	return &BlocBuilder{
		Bloc:    driver,
		Builder: builder,
	}
}

type BlocBuilder struct {
	vecty.Core
	Bloc          bloc.Driver                                  `vecty:"prop"`
	Builder       func(state bloc.State) vecty.ComponentOrHTML `vecty:"prop"`
	OnSubscribe   func(initState bloc.State)                   `vecty:"prop"`
	//OnUnsubscribe func()                                       `vecty:"prop"`
	State         bloc.State                                   `vecty:"prop"`
	StateChan     bloc.Subscription                            `vecty:"prop"`

	prevBloc bloc.Driver
	prevStateChan bloc.Subscription
}

func (p *BlocBuilder) Render() vecty.ComponentOrHTML {
	if p.prevBloc == nil {
		p.prevBloc = p.Bloc
	} else if p.prevBloc != p.Bloc {
		// The component properties have changed but the component type has not.
		// Therefore, Vecty will not call Unmount(), even though the component is configured different.
		// So need to reset this bloc manually.
		if (p.prevStateChan != nil) {
			p.prevStateChan.Close()
		}
		p.prevBloc = p.Bloc
		p.State = nil
	}

	if p.State == nil {
		p.blocState()
	}

	return p.Builder(p.State)
}

func (b *BlocBuilder) blocState() bloc.State {
	if b.StateChan == nil {
		b.StateChan = b.Bloc.Subscribe()
		b.State = <-b.StateChan.Chan()
		b.prevStateChan = b.StateChan

		go func() {
			//runtime.LockOSThread()
			for state := range b.StateChan.Chan() {
				//if state != b.State {
					b.State = state
					vecty.Rerender(b)
				//}
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

	//if b.OnUnsubscribe != nil {
	//	b.OnUnsubscribe()
	//}
}