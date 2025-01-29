package blocks

import (
	"scratcheditor/utils"
	"time"
)

type OnGrabFunc func(e EventGrab)

var (
	onGrabFunc     func(e EventGrab)
	onGrabDebounce float64 //! Implement
)

func OnGrab(f OnGrabFunc, debounce float64) {
	onGrabFunc = f
	onGrabDebounce = debounce
}

func handleOnGrab(b *Block, cursor *utils.Cursor) {
	if b.IsGrabbed {
		return
	}
	pos := cursor.GetPosition()
	if !b.BoundingBox.IsWithin(b.Position, pos) || !cursor.JustPressedLMB() {
		return
	}

	now := time.Now()
	b.IsGrabbed = true
	b.grabWhen = now
	cursor.Grabbed = b.UID
	b.GrabOffset = pos.Sub(b.Position)
	if onGrabFunc == nil {
		return
	}
	onGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  now,
		},
		CursorPosition: pos,
		Offset:         b.GrabOffset,
	})
}
