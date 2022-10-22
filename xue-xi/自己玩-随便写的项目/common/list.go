package common

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type SwapQueue struct {

	InputLen int64

	OutputLen int64

	sync *sync.Mutex

	InputList *list.List

	OutputList *list.List
}

func (q *SwapQueue) NewSwapQueue () *SwapQueue {
	return &SwapQueue{}
}

func(q *SwapQueue) Len() int {
	return int(atomic.LoadInt64(&q.OutputLen))
}

func (q *SwapQueue) Offer(v interface{}) {

	q.sync.Lock()
	defer q.sync.Unlock()

	q.InputList.PushBack(v)

	atomic.AddInt64(&q.InputLen , 1)
}

func (q *SwapQueue) SwapAndConSum () *list.List {
	q.Swap()
	return q.ConSum()

}

func (q *SwapQueue) Swap() {

	q.sync.Lock()
	defer q.sync.Unlock()

	q.InputList  , q.OutputList = q.OutputList ,q.InputList
	q.InputLen , q.OutputLen = q.OutputLen ,q.InputLen

}

func (q *SwapQueue) ConSum() *list.List {
	q.sync.Lock()
	defer q.sync.Unlock()
	list := new(list.List)

	if q.Len() > 0 {
		front := q.OutputList.Front()
		list.PushBack(front)
		q.Dequeue()
	}
	return list
}

func(q *SwapQueue) Dequeue() {
	atomic.AddInt64(&q.OutputLen , -1)
}
