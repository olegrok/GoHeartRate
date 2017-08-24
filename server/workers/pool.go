package workers

import (
	"sync"
	"time"

	"errors"
	"github.com/olegrok/GoHeartRate/protocol"
)

type function func() interface{}

type task struct {
	f      function
	wg     sync.WaitGroup
	result interface{}
}

// Pool is a pool of GoRoutines
type Pool struct {
	concurrency int
	tasksChan   chan *task
	wg          sync.WaitGroup
}

// NewPool is a function that creates new pool of GoRoutines with "concurrency" size
func NewPool(concurrency int) *Pool {
	return &Pool{
		concurrency: concurrency,
		tasksChan:   make(chan *task, concurrency),
	}
}

// Run starts pool of GoRoutines
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		p.wg.Add(1)
		go p.runWorker()
	}
}

// Stop closes channel of tasks and stops all GoRoutines of pool
func (p *Pool) Stop() {
	close(p.tasksChan)
	p.wg.Wait()
}

// AddTaskSync adds new task in pool's queue
func (p *Pool) AddTaskSync(f function) interface{} {
	t := task{
		f:  f,
		wg: sync.WaitGroup{},
	}
	t.wg.Add(1)
	p.tasksChan <- &t
	t.wg.Wait()
	return t.result
}

// AddTaskSyncTimed adds new task in pool's queue with timeout.
// If the task is not taken from the queue during the timeout, then the function returns an error
func (p *Pool) AddTaskSyncTimed(f function, timeout time.Duration) (interface{}, error) {
	t := task{
		f:  f,
		wg: sync.WaitGroup{},
	}
	t.wg.Add(1)

	select {
	case p.tasksChan <- &t:
		break
	case <-time.After(timeout):
		return nil, errors.New(protocol.ErrJobTimedOut)
	}
	t.wg.Wait()
	return t.result, nil
}

func (p *Pool) runWorker() {
	for t := range p.tasksChan {
		t.result = t.f()
		t.wg.Done()
	}
	p.wg.Done()
}
