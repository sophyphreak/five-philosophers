package main

type philosopher struct {
	name  int
	left  chan struct{}
	right chan struct{}
}
