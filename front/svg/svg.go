package svg

import (
	"encoding/xml"
	"image/color"
)

// Main structure
type SVG struct {
	Width       float64
	Height      float64
	Definitions map[string]interface{}
	Elements    []ElementData
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

type ElementData interface {
	ID() string
	Class() string
	Style() []Style
	Fill() color.RGBA
	Stroke() color.RGBA
	StrokeWidth() float32
	Opacity() string
	Visibility() string
}

type Element struct {
	_ID          string
	_Class       string
	_Style       []Style
	_Fill        color.RGBA
	_Stroke      color.RGBA
	_StrokeWidth float32
	_Opacity     string
	_Visibility  string
}

func (e *Element) ID() string {
	return e._ID
}

func (e *Element) Class() string {
	return e._Class
}

func (e *Element) Style() []Style {
	return e._Style
}

func (e *Element) Fill() color.RGBA {
	return e._Fill
}

func (e *Element) Stroke() color.RGBA {
	return e._Stroke
}

func (e *Element) StrokeWidth() float32 {
	return e._StrokeWidth
}

func (e *Element) Opacity() string {
	return e._Opacity
}

func (e *Element) Visibility() string {
	return e._Visibility
}

type RawElementData interface {
	ID() string
	Class() string
	Style() string
	Fill() string
	Stroke() string
	StrokeWidth() string
	Opacity() string
	Visibility() string
}

type RawElement struct {
	ID_          string `xml:"id,attr,omitempty"`
	Class_       string `xml:"class,attr,omitempty"`
	Style_       string `xml:"style,attr,omitempty"`
	Fill_        string `xml:"fill,attr,omitempty"`
	Stroke_      string `xml:"stroke,attr,omitempty"`
	StrokeWidth_ string `xml:"stroke-width,attr,omitempty"`
	Opacity_     string `xml:"opacity,attr,omitempty"`
	Visibility_  string `xml:"visibility,attr,omitempty"`
}

func (e *RawElement) ID() string {
	return e.ID_
}

func (e *RawElement) Class() string {
	return e.Class_
}

func (e *RawElement) Style() string {
	return e.Style_
}

func (e *RawElement) Fill() string {
	return e.Fill_
}

func (e *RawElement) Stroke() string {
	return e.Stroke_
}

func (e *RawElement) StrokeWidth() string {
	return e.StrokeWidth_
}

func (e *RawElement) Opacity() string {
	return e.Opacity_
}

func (e *RawElement) Visibility() string {
	return e.Visibility_
}

type RawCircleElement struct {
	RawElement
	CX string `xml:"cx,attr"`
	CY string `xml:"cy,attr"`
	R  string `xml:"r,attr"`
}

type RawEllipseElement struct {
	RawElement
	CX string `xml:"cx,attr"`
	CY string `xml:"cy,attr"`
	RX string `xml:"rx,attr"`
	RY string `xml:"ry,attr"`
}

type RawLineElement struct {
	RawElement
	X1 string `xml:"x1,attr"`
	Y1 string `xml:"y1,attr"`
	X2 string `xml:"x2,attr"`
	Y2 string `xml:"y2,attr"`
}

type RawPolygonElement struct {
	RawElement
	Points string `xml:"points,attr"`
}

type RawPolylineElement struct {
	RawElement
	Points string `xml:"points,attr"`
}
