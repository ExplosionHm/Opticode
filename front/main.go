package main

import (
	"fmt"
	"log"
	"scratcheditor/blocks"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Block *blocks.Block
}

func (g *Game) Update() error {
	blocks.Handle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//x, y := ebiten.CursorPosition()

	//g.Block.Position.Set(float64(x), float64(y))

	blocks.Render(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Block: X: %.2f, Y: %.2f | %t", g.Block.Position.X, g.Block.Position.Y, g.Block.IsGrabbed))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	blocks.OnGrab(func(e blocks.EventGrab) {
		log.Println("Grabbed")
	}, 0.1)

	blocks.WhileGrab(func(e blocks.EventGrab) {
		log.Println("Grabbing")
		e.Block.Position.Set(e.Cursor.X, e.Cursor.Y)
	}, 0.1)

	blocks.OffGrab(func(e blocks.EventGrab) {
		log.Println("No longer Grabbing")
		e.Block.Position.Set(0, 0)
	}, 0.1)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		Block: blocks.NewBlock(0, 0xff4fd3),
	}); err != nil {
		log.Fatal(err)
	}
}
