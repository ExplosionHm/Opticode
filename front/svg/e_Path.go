package svg

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"regexp"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

type RawPathElement struct {
	RawElement
	D          string `xml:"d,attr"`
	PathLength string `xml:"pathLength,attr"`
}

func (rpe *RawPathElement) Parse() (*PathElement, error) {
	pe := &PathElement{}

	commandRegex := regexp.MustCompile(`(?i)([a-zA-Z])|(-?\d*\.?\d+(?:e[-+]?\d+)?)`)
	matches := commandRegex.FindAllString(rpe.D, -1)

	if len(matches) == 0 {
		return nil, fmt.Errorf("no path data found")
	}
	var currentCommand rune
	var currentParams []float64

	for _, match := range matches {
		if len(match) == 1 && ((match[0] >= 'A' && match[0] <= 'Z') || (match[0] >= 'a' && match[0] <= 'z')) {
			if currentCommand != 0 {
				pe.D = append(pe.D, PathCommand{Type: currentCommand, Params: currentParams})
				currentParams = nil
			}

			currentCommand = rune(match[0])
		} else {

			param, err := strconv.ParseFloat(match, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid parameter '%s': %v", match, err)
			}
			currentParams = append(currentParams, param)
		}
	}

	if currentCommand != 0 {
		pe.D = append(pe.D, PathCommand{Type: currentCommand, Params: currentParams})
	}
	if len(rpe.PathLength) > 0 {
		var err error
		pe.PathLength, err = strconv.Atoi(rpe.PathLength)
		if err != nil {
			return nil, fmt.Errorf("invalid pathLength %v", err)
		}
	} else {
		if pe.PathLength == 0 {
			pe.PathLength = -1
		}
	}

	pe._Fill = GetColorFromString(rpe.Fill_)
	pe._Stroke = GetColorFromString(rpe.Stroke_)
	if rpe.StrokeWidth_ == "" {
		sw, err := strconv.ParseFloat(rpe.StrokeWidth_, 32)
		if err != nil {
			return nil, err
		}
		pe._StrokeWidth = float32(sw)
	}

	return pe, nil
}

type PathElement struct {
	Element
	D          []PathCommand
	PathLength int
}

type PathCommand struct {
	Type   rune
	Params []float64
}

func (pe *PathElement) Draw(dst *ebiten.Image, defs map[string]interface{}) error {
	whiteImage.Fill(color.White)
	var x, y float64
	var startX, startY float64

	var prevCurveType rune
	var prevCtrlX float64
	var prevCtrlY float64
	var currX float64
	var currY float64

	path := vector.Path{}
	for _, cmd := range pe.D {
		switch cmd.Type {
		case 'M', 'm': // Move to
			x, y = cmd.Params[0], cmd.Params[1]
			if cmd.Type == 'm' {
				x += startX
				y += startY
			}
			startX, startY = x, y
			path.MoveTo(float32(x), float32(y))
		case 'L', 'l': // Line to
			nx, ny := cmd.Params[0], cmd.Params[1]
			if cmd.Type == 'l' {
				nx += x
				ny += y
			}
			path.LineTo(float32(nx), float32(ny))
			x, y = nx, ny
		case 'C', 'c': // Cubic Bezier Curve
			cx1, cy1 := cmd.Params[0], cmd.Params[1]
			cx2, cy2 := cmd.Params[2], cmd.Params[3]
			nx, ny := cmd.Params[4], cmd.Params[5]
			if cmd.Type == 'c' {
				cx1 += x
				cy1 += y
				cx2 += x
				cy2 += y
				nx += x
				ny += y
			}
			path.CubicTo(float32(cx1), float32(cy1), float32(cx2), float32(cy2), float32(nx), float32(ny))
			x, y = nx, ny
			prevCurveType = 'C'
			prevCtrlX = cmd.Params[2]
			prevCtrlY = cmd.Params[3]
			currX = cmd.Params[4]
			currY = cmd.Params[5]

		case 'S', 's': // Smooth cubic Bezier curve to (continuation of previous cubic curve)
			var cx1, cy1 float64

			if prevCurveType == 'C' || prevCurveType == 'c' {
				cx1 = currX + (currX - prevCtrlX)
				cy1 = currY + (currY - prevCtrlY)
			} else {
				cx1 = currX
				cy1 = currY
			}

			cx2, cy2 := cmd.Params[0], cmd.Params[1]
			nx, ny := cmd.Params[2], cmd.Params[3]

			if cmd.Type == 's' {
				cx2 += x
				cy2 += y
				nx += x
				ny += y
			}

			path.CubicTo(float32(cx1), float32(cy1), float32(cx2), float32(cy2), float32(nx), float32(ny))
			prevCtrlX = cx2
			prevCtrlY = cy2
			x, y = nx, ny

			/*
				case 'Q', 'q': // Quadratic Bezier Curve
					for i := 0; i < len(cmd.Params); i += 4 {
						cx, cy := cmd.Params[i], cmd.Params[i+1]
						nx, ny := cmd.Params[i+2], cmd.Params[i+3]
						if cmd.Type == 'q' {
							cx += x
							cy += y
							nx += x
							ny += y
						}
						vertices, indices = addQuadratic(vertices, indices, x, y, cx, cy, nx, ny, 1.5)
						x, y = nx, ny
					}
						case 'T', 't': // Smooth quadratic Bezier curve to
							for i := 0; i < len(cmd.Params); i += 2 {
								cx, cy := 2*x-x, 2*y-y // Reflect previous control point
								nx, ny := cmd.Params[i], cmd.Params[i+1]
								if cmd.Type == 't' {
									nx += x
									ny += y
								}
								vertices, indices = addQuadratic(vertices, indices, x, y, cx, cy, nx, ny, 1.5)
								x, y = nx, ny
							}
						case 'A', 'a': // Arc
							for i := 0; i < len(cmd.Params); i += 7 {
								rx, ry := cmd.Params[i], cmd.Params[i+1]
								rotation := cmd.Params[i+2]
								largeArc := cmd.Params[i+3] != 0
								sweep := cmd.Params[i+4] != 0
								nx, ny := cmd.Params[i+5], cmd.Params[i+6]
								if cmd.Type == 'a' {
									nx += x
									ny += y
								}
								vertices, indices = addArc(vertices, indices, x, y, rx, ry, rotation, largeArc, sweep, nx, ny, 1.5)
								x, y = nx, ny
							}
						case 'Z', 'z': // Close path
							vertices, indices = addLine(vertices, indices, float32(x), float32(y), float32(startX), float32(startY), 1.5)
							x, y = startX, startY
			*/
		default:
			log.Printf("unknown type %c %+v", cmd.Type, cmd)
		}
	}

	op := &vector.StrokeOptions{}
	op.Width = pe._StrokeWidth
	op.LineJoin = vector.LineJoinRound //! Implement svg defintion
	svertices, sindices := path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	for i := range vertices {
		vertices[i].DstX = (vertices[i].DstX + float32(x))
		vertices[i].DstY = (vertices[i].DstY + float32(y))
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(pe._Fill.R)
		vertices[i].ColorG = float32(pe._Fill.G)
		vertices[i].ColorB = float32(pe._Fill.B)
		vertices[i].ColorA = float32(pe._Fill.A)
	}

	for i := range svertices {
		svertices[i].DstX = (svertices[i].DstX + float32(x))
		svertices[i].DstY = (svertices[i].DstY + float32(y))
		svertices[i].SrcX = 1
		svertices[i].SrcY = 1
		svertices[i].ColorR = float32(pe._Stroke.R)
		svertices[i].ColorG = float32(pe._Stroke.G)
		svertices[i].ColorB = float32(pe._Stroke.B)
		svertices[i].ColorA = float32(pe._Stroke.A)
	}

	opt := &ebiten.DrawTrianglesOptions{}
	opt.AntiAlias = true
	opt.FillRule = ebiten.FillRuleNonZero

	dst.DrawTriangles(vertices, indices, whiteSubImage, opt)
	dst.DrawTriangles(svertices, sindices, whiteSubImage, opt)

	return nil
}
