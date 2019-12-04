package bloc

import "container/list"

type subscriptionList struct {
	list	*list.List
}

func newSubscriptionList() *subscriptionList {
	return &subscriptionList{list.New()}
}

func (sl *subscriptionList) add(state State) *listSubscription {
	sub := &listSubscription{ list: sl, C: make(chan State, 1)}
	sub.e = sl.list.PushBack(sub)
	sub.C <- state
	return sub
}

func (sl *subscriptionList) remove(s *listSubscription) {
	sl.list.Remove(s.e)
}

func (sl *subscriptionList) notify(state State) {
	for e := sl.list.Front(); e != nil; e = e.Next() {
		e.Value.(*listSubscription).C <- state
	}
}

func (sl *subscriptionList) close() {
	for e := sl.list.Front(); e != nil; e = e.Next() {
		if e.Value.(*listSubscription).C != nil {
			close(e.Value.(*listSubscription).C)
			e.Value.(*listSubscription).C = nil
		}
		//e.Value.(*listSubscription).Close()
	}
	sl.list = nil
}


type listSubscription struct {
	C       chan State
	list	*subscriptionList
	e		*list.Element
}

func (s *listSubscription) Chan() chan State {
	return s.C
}

func (s *listSubscription) Close() {
	close(s.C)
	s.C = nil
	s.list.remove(s)
}
