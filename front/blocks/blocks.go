package blocks

import (
	"image/color"
	"log"
	"math/rand"
	"scratcheditor/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	MaxBlocks = 9000
)

var ()

type Block struct {
	Type        uint8
	Color       uint32
	UID         uint32
	Position    utils.Vector
	BoundingBox utils.Box

	Grabbable  bool
	IsGrabbed  bool
	GrabOffset utils.Vector
	grabWhen   time.Time
}

type BlockDrawOptions interface {
}

// Create an array to store all blocks
var Blocks = make([]*Block, 0, MaxBlocks)

func appendBlock(b *Block) {
	if len(Blocks) > MaxBlocks {
		// Add proper handling
		log.Fatal("Block limit reached!")
	}
	Blocks = append(Blocks, b)
}

func NewBlock(t uint8, color uint32) *Block {
	b := &Block{
		Type:        t,
		Color:       color,
		Position:    utils.Vector{X: 0, Y: 0},
		Grabbable:   true,
		UID:         rand.Uint32(),
		BoundingBox: utils.NewBox(0, 0, 100, 50),
	}
	appendBlock(b)
	return b
}

func HexToRGBA(hex uint32) color.RGBA {
	return color.RGBA{
		R: uint8((hex >> 16) & 0xFF),
		G: uint8((hex >> 8) & 0xFF),
		B: uint8((hex) & 0xFFb),
		A: 255,
	}
}

func (b *Block) Draw(image *ebiten.Image, op BlockDrawOptions) {
	ebitenutil.DrawRect(image, b.Position.X, b.Position.Y, 100, 50, HexToRGBA(b.Color))
}

func Render(image *ebiten.Image) {
	/* for i := 0; i < len(Blocks); i++ {
		Blocks[i].Draw(image, nil)
	} */

	for i := len(Blocks) - 1; i >= 0; i-- {
		Blocks[i].Draw(image, nil)
	}
}
