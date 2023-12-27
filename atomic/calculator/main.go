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
