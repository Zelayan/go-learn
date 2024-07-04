package main

import (
	"sync"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	sync.Mutex
}

func (l *Lock) Write() {
	l.Lock()
	l.count++
	time.Sleep(cost)
	l.Unlock()
}

func (l *Lock) Read() {
	l.Lock()
	time.Sleep(cost)
	_ = l.count
	l.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}
