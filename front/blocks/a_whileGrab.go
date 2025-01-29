package blocks

import (
	"scratcheditor/utils"

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

func handlewhileGrab(b *Block, cursor *utils.Cursor) {
	if !b.IsGrabbed {
		return
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		handleOffGrab(b, cursor)
		return
	}

	whileGrabFunc(EventGrab{
		EventOps: EventOps{
			Block: b,
			When:  b.grabWhen,
		},
		CursorPosition: cursor.GetPosition(),
		Offset:         b.GrabOffset,
	})
}
