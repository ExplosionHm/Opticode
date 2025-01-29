package svg

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"strconv"
)

func LoadFromFile(path string) (*SVG, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.New("file path does not exist.")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	con, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return Load(con)
}

func Load(svg []byte) (*SVG, error) {
	// Parse the XML
	var _xml XML
	err := xml.Unmarshal(svg, &_xml)
	if err != nil {
		return nil, err
	}

	return ParseSVG(_xml)
}

func ParseSVG(_xml XML) (*SVG, error) {
	decoder := xml.NewDecoder(bytes.NewReader([]byte(_xml.InnerXML)))

	h, err := strconv.ParseFloat(_xml.Height, 64)
	if err != nil {
		return nil, err
	}
	w, err := strconv.ParseFloat(_xml.Width, 64)
	if err != nil {
		return nil, err
	}

	elements := []interface{}{}
	definitions := make(map[string]interface{})
	inDefs := false

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch element := token.(type) {
		// TODO: Add error cases
		case xml.StartElement:
			switch element.Name.Local {
			case "defs":
				inDefs = true
			case "rect":
				var rect RectElement
				decoder.DecodeElement(&rect, &element)
				if inDefs {
					definitions[rect.ID] = rect
				} else {
					elements = append(elements, rect)
				}
			case "circle":
				var circle CircleElement
				decoder.DecodeElement(&circle, &element)
				if inDefs {
					definitions[circle.ID] = circle
				} else {
					elements = append(elements, circle)
				}
			case "ellipse":
				var ellipse EllipseElement
				decoder.DecodeElement(&ellipse, &element)
				if inDefs {
					definitions[ellipse.ID] = ellipse
				} else {
					elements = append(elements, ellipse)
				}
			case "line":
				var line LineElement
				decoder.DecodeElement(&line, &element)
				if inDefs {
					definitions[line.ID] = line
				} else {
					elements = append(elements, line)
				}
			case "polygon":
				var polygon PolygonElement
				decoder.DecodeElement(&polygon, &element)
				if inDefs {
					definitions[polygon.ID] = polygon
				} else {
					elements = append(elements, polygon)
				}
			case "polyline":
				var polyline PolylineElement
				decoder.DecodeElement(&polyline, &element)
				if inDefs {
					definitions[polyline.ID] = polyline
				} else {
					elements = append(elements, polyline)
				}
			case "path":
				var path PathElement
				decoder.DecodeElement(&path, &element)
				if inDefs {
					definitions[path.ID] = path
				} else {
					elements = append(elements, path)
				}
			}
		case xml.EndElement:
			if element.Name.Local == "defs" {
				inDefs = false
			}
		}
	}
	return &SVG{
		Height:   h,
		Width:    w,
		Elements: elements,
	}, nil
}
