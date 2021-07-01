package concurrent

import (
	"sync"
)

type WorkerPool struct {
	NumOfWorker int
	quit        chan bool
	process     chan FutureTask
	wg          sync.WaitGroup
	lock        sync.Mutex
}

type FutureTask struct {
	RunFunc func()
	done    bool
}

func NewWorkerPool(numWorker int) *WorkerPool {
	w := &WorkerPool{}
	var p sync.WaitGroup
	var s sync.Mutex
	w.NumOfWorker = numWorker
	w.wg = p
	w.lock = s
	w.process = make(chan FutureTask, numWorker)
	w.quit = make(chan bool, 1)
	return w
}

func NewFutureTask(fun func()) *FutureTask {
	f := &FutureTask{
		RunFunc: fun,
		done:    false,
	}
	return f
}

func (w *WorkerPool) AddTask(f *FutureTask) {
	w.lock.Lock()
	w.wg.Add(1)
	w.lock.Unlock()
	w.process <- *f
}

func (w *WorkerPool) Wait() {
	w.wg.Wait()
	w.quit <- true
}

func (w *WorkerPool) Run() {
	for {
		select {
		case t, ok := <-w.process:
			if !ok {
				break
			}
			go func() {
				t.RunFunc()
				t.done = true
				w.wg.Done()
			}()
		case <-w.quit:
			break
		}
	}
}
