package semaphore

import (
	"sync"
)

type Semaphore struct {
	channel chan int
	wg      sync.WaitGroup
}

func New(max int) Semaphore {
	x := Semaphore{
		channel: make(chan int, max),
	}
	return x
}

func (s *Semaphore) Add() {
	s.wg.Add(1)
	s.channel <- 0
}

func (s *Semaphore) Done() {
	<-s.channel
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
