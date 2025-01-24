package blocks

import (
	"scratcheditor/utils"
	"time"
)

type OffGrabFunc func(e EventGrab)

var (
	offGrabFunc     func(e EventGrab)
	offGrabDebounce float64 //! Implement
)

func OffGrab(f OffGrabFunc, debounce float64) {
	offGrabFunc = f
	offGrabDebounce = debounce
}

func handleOffGrab(b *Block, cursor utils.Vector, when time.Time) {
	b.IsGrabbed = false
	if offGrabFunc == nil {
		return
	}
	offGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  when,
		},
		Cursor: cursor,
	})
}
