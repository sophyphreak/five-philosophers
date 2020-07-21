package main

import (
	"fmt"
	"sync"
)

func main() {
	var forks [5]chan struct{}
	for i := range forks {
		forks[i] = make(chan struct{}, 1)
	}
	var philosophers [5]philosopher
	for i := range philosophers {
		l := i - 1
		r := i
		if l < 0 {
			l = 4
		}
		philosophers[i] = philosopher{i, forks[l], forks[r]}
	}
	fmt.Println(philosophers)
	for i := range philosophers {
		go philosophers[i].eat()
	}
	for _, f := range forks {
		f <- struct{}{}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
