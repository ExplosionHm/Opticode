package svg

type RawRectElement struct {
	RawElement
	X      string `xml:"x,attr"`
	Y      string `xml:"y,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	RX     string `xml:"rx,attr,omitempty"`
	RY     string `xml:"ry,attr,omitempty"`
}
