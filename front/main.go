package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"scratcheditor/blocks"
	"scratcheditor/svg"
	"scratcheditor/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Block  *blocks.Block
	Cursor *utils.Cursor
}

var l = time.Now()

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Since(l).Seconds() > 0.0 {
		cx, cy := ebiten.CursorPosition()
		blocks.NewBlock(0, rand.Uint32()).Position.Set(float64(cx), float64(cy))
	}

	blocks.Handle(g.Cursor)
	g.Cursor.Handle()
	return nil
}

var offset = utils.Vector{}

func (g *Game) Draw(screen *ebiten.Image) {
	blocks.Render(screen)

	cx, cy := ebiten.CursorPosition()

	vector.StrokeLine(screen, 0, 0, float32(cx), float32(cy), 1, color.RGBA{255, 0, 0, 255}, true)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Fps: %.2f\nBlock: X: %.2f, Y: %.2f | %t | Offset: X: %.2f, Y: %.2f\nCursor: %d\nTotal Blocks: %d", ebiten.ActualFPS(), g.Block.Position.X, g.Block.Position.Y, g.Block.IsGrabbed, offset.X, offset.Y, g.Cursor.Grabbed, len(blocks.Blocks)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	p := time.Now()
	s, err := svg.LoadFromFile("stack.svg")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%d microseconds", time.Since(p).Microseconds())
	for i, e := range s.Elements {
		if e == nil {
			log.Println(i, "is nil")
			continue
		}
		switch ele := e.(type) {
		case svg.RectElement:
			log.Println(ele)
		case svg.PolylineElement:
			log.Println(ele)
		case svg.PolygonElement:
			log.Println(ele)
		case svg.PathElement:
			log.Println(ele)
		case svg.LineElement:
			log.Println(ele)
		case svg.EllipseElement:
			log.Println(ele)
		case svg.CircleElement:
			log.Println(ele)
		default:
			log.Println("UNKNOWN:", e)
		}
	}
	blocks.OnGrab(func(e blocks.EventGrab) {
		log.Println("Grabbed")
		offset = e.Offset
		e.Block.Color = 0x4ceb34
	}, 0.1)

	blocks.WhileGrab(func(e blocks.EventGrab) {
		log.Println("Grabbing", time.Since(e.When).Seconds(), e.Offset)
		e.Block.Position.Set(e.CursorPosition.X-e.Offset.X, e.CursorPosition.Y-e.Offset.Y)
	}, 0.1)

	blocks.OffGrab(func(e blocks.EventGrab) {
		log.Println("No longer Grabbing")
		e.Block.Color = 0xeb3449
		e.Block.Grabbable = true
		//e.Block.Position.Set(0, 0)
	}, 0.1)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GoScratch")
	if err := ebiten.RunGame(&Game{
		Block:  blocks.NewBlock(0, 0xff4fd3),
		Cursor: &utils.Cursor{},
	}); err != nil {
		log.Fatal(err)
	}
}
