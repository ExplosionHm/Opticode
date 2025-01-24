package blocks

import (
	"scratcheditor/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Define base struct for other events to in inherient

type EventOps struct {
	Block *Block
	When  time.Time
}

// Event Grab parameters
type EventGrab struct {
	EventOps
	Cursor utils.Vector
	DeltaX float64 //! Implement
	DeltaY float64 //! Implement
}

var (
	grabWhen time.Time
)

func Handle() {
	cX, cY := ebiten.CursorPosition()
	cursor := utils.Vector{X: float64(cX), Y: float64(cY)}

	for i := 0; i < len(Blocks); i++ {
		b := Blocks[i]
		if t := handleOnGrab(b, cursor); t != nil {
			grabWhen = *t
		}
		handlewhileGrab(b, cursor, grabWhen)
	}
}
