package utils

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	LMBForgiveness = 0.1
)

type Cursor struct {
	Grabbed    uint32
	LMB        time.Time
	lmbPressed bool
	RMB        time.Time
	rmbPressed bool
}

func (c *Cursor) GetPosition() Vector {
	x, y := ebiten.CursorPosition()
	return Vector{X: float64(x), Y: float64(y)}
}

func (c *Cursor) JustPressedLMB() bool {
	return time.Since(c.LMB).Seconds() < LMBForgiveness
}

func (c *Cursor) Handle() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !c.lmbPressed {
			c.LMB = time.Now()
			c.lmbPressed = true
		}
	} else {
		if c.lmbPressed {
			c.lmbPressed = false
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if !c.rmbPressed {
			c.RMB = time.Now()
			c.rmbPressed = true
		}
	} else {
		if c.rmbPressed {
			c.rmbPressed = false
		}
	}
}
