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

type Editor struct {
	Block  *blocks.Block
	Cursor *utils.Cursor
	Test   *svg.SVG
}

var l = time.Now()

func (g *Editor) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Since(l).Seconds() > 0.0 {
		cx, cy := ebiten.CursorPosition()
		blocks.NewBlock(0, rand.Uint32()).Position.Set(float64(cx), float64(cy))
	}

	blocks.Handle(g.Cursor)
	g.Cursor.Handle()
	return nil
}

var offset = utils.Vector{}

func (g *Editor) Draw(screen *ebiten.Image) {
	blocks.Render(screen)

	cx, cy := ebiten.CursorPosition()

	vector.StrokeLine(screen, 0, 0, float32(cx), float32(cy), 1, color.RGBA{255, 0, 0, 255}, true)
	g.Test.Draw(screen, utils.Vector{X: float64(cx), Y: float64(cy)})

	debug := fmt.Sprintf("Fps: %.2f\nBlock: X: %.2f, Y: %.2f | %t | Offset: X: %.2f, Y: %.2f\nCursor: %d\nTotal Blocks: %d", ebiten.ActualFPS(), g.Block.Position.X, g.Block.Position.Y, g.Block.IsGrabbed, offset.X, offset.Y, g.Cursor.Grabbed, len(blocks.Blocks))
	_, y := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, debug, 0, y-100)
}

func (g *Editor) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	/* p := time.Now()
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
	} */
	/* re := `<svg width="300" height="130" xmlns="http://www.w3.org/2000/svg">
	  <path
	    fill="red"
	    stroke="blue"
	    d="M 10,90
	           C 30,90 25,10 50,10
	           S 70,90 90,90" />
	</svg>
	` */
	//<rect width="200" height="100" x="10" y="10" rx="20" ry="20" fill="blue" />
	s, err := svg.LoadFromFile("svgtest.svg")
	if err != nil {
		log.Println(err)
	}

	blocks.OnGrab(func(e blocks.EventGrab) {
		log.Println("Grabbed")
		offset = e.Offset
		e.Block.Color = 0x4ceb34
	}, 0.1)

	blocks.WhileGrab(func(e blocks.EventGrab) {
		//log.Println("Grabbing", time.Since(e.When).Seconds(), e.Offset)
		e.Block.Position.Set(e.CursorPosition.X-e.Offset.X, e.CursorPosition.Y-e.Offset.Y)
		log.Println(e.CursorPosition)
	}, 0.1)

	blocks.OffGrab(func(e blocks.EventGrab) {
		log.Println("No longer Grabbing")
		e.Block.Color = 0xeb3449
		//e.Block.Position.Set(0, 0)
	}, 0.1)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GoScratch")
	if err := ebiten.RunGame(&Editor{
		Block:  blocks.NewBlock(0, 0xff4fd3),
		Cursor: &utils.Cursor{},
		Test:   s,
	}); err != nil {
		log.Fatal(err)
	}
}
