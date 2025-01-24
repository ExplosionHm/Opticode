package blocks

import (
	"scratcheditor/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type WhileGrabFunc func(e EventGrab)

var (
	whileGrabFunc     func(e EventGrab)
	whileGrabDebounce float64 //! Implement
)

func WhileGrab(f WhileGrabFunc, debounce float64) {
	whileGrabFunc = f
	whileGrabDebounce = debounce
}

func handlewhileGrab(b *Block, cursor utils.Vector, when time.Time) {
	if !b.IsGrabbed {
		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		handleOffGrab(b, cursor, when)
		return
	}

	whileGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  when,
		},
		Cursor: cursor,
	})
}
