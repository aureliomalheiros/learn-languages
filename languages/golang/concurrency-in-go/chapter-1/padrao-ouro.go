package main

import (
	"fmt"
)

//ch chan func(int) int
//    |    |
//    |    └─ tipo: uma função que recebe int e retorna int
//    └─────── é um canal deste tipo

type SafeCounter struct {
	ch chan func(int) int
	val int
}

func NewSafeCounter() *SafeCounter {
	sc := &SafeCounter{
		ch: make(chan func(int) int),
	}
	go sc.loop()
	return sc
}

func (sc *SafeCounter) loop() {
	for op := range sc.ch {
		sc.val = op(sc.val)
	}
}

func (sc *SafeCounter) Add(delta int) {
	sc.ch <- func(old int) int { return old + delta }
}

func (sc *SafeCounter) Value() int {
	result := make(chan int)
	sc.ch <- func(old int) int {
		result <- old
		return old
	}
	return <-result
}

func main() {
	counter := NewSafeCounter()
	counter.Add(5)
	counter.Add(3)
	fmt.Println(counter.val)
}
