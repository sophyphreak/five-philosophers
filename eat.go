package main

import (
	"fmt"
	"time"
)

func (p *philosopher) eat() {
	leftHand, rightHand := false, false
GetFork:
	select {
	case <-p.left:
		if p.name == 0 {
			fmt.Println("Philosopher", p.name, "picks up left fork at", 4)
		} else {
			fmt.Println("Philosopher", p.name, "picks up left fork at", p.name-1)
		}
		leftHand = true
	case <-p.right:
		fmt.Println("Philosopher", p.name, "picks up right fork at", p.name)
		rightHand = true
	}
	if !(leftHand && rightHand) {
		goto GetFork
	}
	fmt.Println("Philosopher", p.name, "is eating.")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Philosopher", p.name, "is thinking.")
	p.left <- struct{}{}
	p.right <- struct{}{}
	leftHand, rightHand = false, false
	time.Sleep(5 * time.Millisecond)
	goto GetFork
}
