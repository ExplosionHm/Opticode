package blocks

import (
	"scratcheditor/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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

func handleOnGrab(b *Block, cursor utils.Vector) *time.Time {
	if b.IsGrabbed {
		return nil
	}

	if !b.BoundingBox.IsWithin(cursor) || !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	b.IsGrabbed = true
	if onGrabFunc == nil {
		return nil
	}
	now := time.Now()
	onGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  now,
		},
		Cursor: cursor,
	})
	return &now
}
