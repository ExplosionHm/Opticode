package svg

import (
	"hash/fnv"
	"image/color"
	"log"
	"scratcheditor/utils"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	colors = map[string]color.RGBA{
		"blue": {0, 0, 255, 255},
		"red":  {255, 0, 0, 255},
	}
)

func GetColorFromString(s string) color.RGBA {
	if c, ok := colors[s]; ok {
		return c
	}
	// convert unknown string to color

	n := strings.ToLower(strings.ReplaceAll(s, " ", ""))

	h := fnv.New32a()
	h.Write([]byte(n))
	hash := h.Sum32()
	log.Println(color.RGBA{
		R: uint8((hash >> 16) & 0xFF),
		G: uint8((hash >> 8) & 0xFF),
		B: uint8(hash & 0xFF),
		A: 255,
	})
	return color.RGBA{
		R: uint8((hash >> 16) & 0xFF),
		G: uint8((hash >> 8) & 0xFF),
		B: uint8(hash & 0xFF),
		A: 255,
	}
}

func (s *SVG) Draw(screen *ebiten.Image, vec utils.Vector) {
	for i, e := range s.Elements {
		if e == nil {
			log.Println(i, "is nil")
			continue
		}
		switch ele := e.(type) {
		case *PathElement:
			ele.Draw(screen, s.Definitions)
		/* case PolylineElement:
			log.Println(ele)
		case PolygonElement:
			log.Println(ele)
		case PathElement:
			log.Println(ele)
		case LineElement:
			log.Println(ele)
		case EllipseElement:
			log.Println(ele)
		case CircleElement:
			log.Println(ele) */
		default:
			log.Println("UNKNOWN:", e)
		}
	}
}
