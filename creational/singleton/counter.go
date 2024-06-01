package main

import "sync"

type Singleton struct {
	counter int
}

var Instance *Singleton

var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		Instance = &Singleton{}
	})

	return Instance
}

func (s *Singleton) Increase() {
	if Instance == nil {
		return
	}
	s.counter++
}

func (s *Singleton) GetCounter() int {
	return s.counter
}
