package svg

import (
	"encoding/xml"
	"image/color"
)

type SVG struct {
	Width       float64
	Height      float64
	Definitions map[string]interface{}
	Elements    []interface{}
}

type Style struct {
	Method     uint16
	Parameters []interface{}
}

type XML struct {
	XMLName  xml.Name `xml:"svg"`
	Width    string   `xml:"width,attr"`
	Height   string   `xml:"height,attr"`
	InnerXML []byte   `xml:",innerxml"`
}

type PElement struct {
	ID         string
	Class      string
	Style      []Style
	Fill       color.Color
	Stroke     string
	Opacity    string
	Visibility string
}

type Element struct {
	ID         string `xml:"id,attr,omitempty"`
	Class      string `xml:"class,attr,omitempty"`
	Style      string `xml:"style,attr,omitempty"`
	Fill       string `xml:"fill,attr,omitempty"`
	Stroke     string `xml:"stroke,attr,omitempty"`
	Opacity    string `xml:"opacity,attr,omitempty"`
	Visibility string `xml:"visibility,attr,omitempty"`
}

type PathElement struct {
	Element
	D string `xml:"d,attr"`
}

type CircleElement struct {
	Element
	CX string `xml:"cx,attr"`
	CY string `xml:"cy,attr"`
	R  string `xml:"r,attr"`
}

type EllipseElement struct {
	Element
	CX string `xml:"cx,attr"`
	CY string `xml:"cy,attr"`
	RX string `xml:"rx,attr"`
	RY string `xml:"ry,attr"`
}

type LineElement struct {
	Element
	X1 string `xml:"x1,attr"`
	Y1 string `xml:"y1,attr"`
	X2 string `xml:"x2,attr"`
	Y2 string `xml:"y2,attr"`
}

type PolygonElement struct {
	Element
	Points string `xml:"points,attr"`
}

type PolylineElement struct {
	Element
	Points string `xml:"points,attr"`
}
