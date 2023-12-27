package main

import (
	"sync/atomic"
)

type calculator struct {
	res atomic.Value
}

func newCalculator() calculator {
	c := calculator{}
	c.res.Store(float64(0))
	return c
}

func (c *calculator) result() float64 {
	r, ok := c.res.Load().(float64)
	if !ok {
		panic("operating with wrong type")
	}
	return r
}

func (c *calculator) add(n float64) {
	c.res.Store(c.result() + n)
}

func (c *calculator) sub(n float64) {
	c.res.Store(c.result() - n)
}

func (c *calculator) mul(n float64) {
	c.res.Store(c.result() * n)
}

func (c *calculator) div(n float64) {
	if n == 0 {
		panic("cannot divide by zero")
	}
	c.res.Store(c.result() / n)
}
