package main

import (
	"sync/atomic"
)

type calculator struct {
	res atomic.Value
}
