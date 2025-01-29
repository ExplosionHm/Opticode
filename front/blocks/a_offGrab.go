package blocks

import (
	"scratcheditor/utils"
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

func handleOffGrab(b *Block, cursor *utils.Cursor) {
	b.IsGrabbed = false
	if offGrabFunc == nil {
		return
	}
	cursor.Grabbed = 0
	offGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  b.grabWhen,
		},
		CursorPosition: cursor.GetPosition(),
	})
}
