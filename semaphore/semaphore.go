package semaphore

import (
	"sync"
)

type Semaphore struct {
	channel chan int
	wg      sync.WaitGroup
	//m sync.Mutex
}

func New(max int) Semaphore {
	x := Semaphore{
		channel: make(chan int, max),
	}
	return x
}

func (s *Semaphore) Add() {
	//s.m.Lock()
	s.wg.Add(1)
	s.channel <- 0
	//s.m.Unlock()
}

func (s *Semaphore) Done() {
	<-s.channel
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
