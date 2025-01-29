package blocks

import (
	"scratcheditor/utils"
	"time"
)

// Define base struct for other events to in inherient

type EventOps struct {
	Block *Block
	When  time.Time
}

// Event Grab parameters
type EventGrab struct {
	EventOps
	CursorPosition utils.Vector
	Offset         utils.Vector
	DeltaX         float64 //! Implement
	DeltaY         float64 //! Implement
}

// Loops through all the blocks in the scene and handles all the events
func Handle(csr *utils.Cursor) {
	for i := 0; i < len(Blocks); i++ {
		b := Blocks[i]
		// < ------- Handle Grabbing
		if b.Grabbable {
			if csr.Grabbed == 0 {
				handleOnGrab(b, csr)
			}
			if csr.Grabbed == b.UID {
				handlewhileGrab(b, csr)
			}
		}
		// > ------- End of Handle Grabbing

	}
}
