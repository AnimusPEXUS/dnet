package main

import (
	"sync"

	"github.com/gotk3/gotk3/glib"
)

func CallGlib(fu func()) {

	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()

	glib.IdleAdd(
		func() {
			cond.L.Lock()
			defer cond.L.Unlock()

			fu()

			cond.Signal()
		},
	)

	cond.Wait()
}
