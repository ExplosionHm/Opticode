package svg

import (
	"image/color"
	"log"
	"scratcheditor/utils"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	colors = map[string]color.Color{
		"blue": color.RGBA{0, 0, 255, 255},
	}
)

func (s *SVG) Draw(screen *ebiten.Image, vec utils.Vector) {
	for i, e := range s.Elements {
		if e == nil {
			log.Println(i, "is nil")
			continue
		}
		switch ele := e.(type) {
		case RectElement:
			x, err := strconv.ParseFloat(ele.X, 64)
			if err != nil {
				log.Println("ERROR CONVERTING", x)
				return
			}
			y, err := strconv.ParseFloat(ele.Y, 64)
			if err != nil {
				log.Println("ERROR CONVERTING", y)
				return
			}
			w, err := strconv.ParseFloat(ele.Width, 64)
			if err != nil {
				log.Println("ERROR CONVERTING", w)
				return
			}
			h, err := strconv.ParseFloat(ele.Height, 64)
			if err != nil {
				log.Println("ERROR CONVERTING", h)
				return
			}
			ebitenutil.DrawRect(screen, vec.X+x, vec.Y+y, w, h, colors[ele.Fill])
		case PolylineElement:
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
			log.Println(ele)
		default:
			log.Println("UNKNOWN:", e)
		}
	}
}
