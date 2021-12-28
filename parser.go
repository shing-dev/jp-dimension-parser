package jp_dimension_parser

import (
	"fmt"
	"regexp"
	"strconv"
)

var numRegex = regexp.MustCompile(`[0-9]+`)

// Dimension represents the object's width x depth x height
type Dimension struct {
	Width  Length
	Depth  Length
	Height Length
}

// Parse returns parsed dimension
// When none of the lengths are parsed, it returns nil
func Parse(s string) *Dimension {
	s = analyze(s)
	dim := Dimension{
		Width:  parseWidth(s),
		Depth:  parseDepth(s),
		Height: parseHeight(s),
	}
	if dim.Width > 0 || dim.Depth > 0 || dim.Height > 0 {
		return &dim
	}
	return nil
}

func parseWidth(s string) Length {
	widths := []string{"幅", "width", "w", "W"}
	for _, w := range widths {
		if length := parseDimension(w, s); length > 0 {
			return length
		}
	}
	return 0
}

func parseDepth(s string) Length {
	depths := []string{"奥行き", "奥行", "depth", "D", "d"}
	for _, d := range depths {
		if length := parseDimension(d, s); length > 0 {
			return length
		}
	}
	return 0
}

func parseHeight(s string) Length {
	heights := []string{"高さ", "高", "height", "H", "h"}
	for _, h := range heights {
		length := parseDimension(h, s)
		if length > 0 {
			return length
		}
	}
	return 0
}

func parseDimension(dimensionName string, s string) Length {
	lengthFormats := []struct {
		format string
		length Length
	}{
		{format: "mm", length: Millimeter},
		{format: "cm", length: Centimeter},
		{format: "m", length: Meter},
	}
	for _, lf := range lengthFormats {
		ranges := []string{
			fmt.Sprintf(`[0-9]+%s`, lf.format),
			fmt.Sprintf(`[0-9]+〜[0-9]+%s`, lf.format),
			fmt.Sprintf(`[0-9]+~[0-9]+%s`, lf.format),
		}
		for _, r := range ranges {
			re := regexp.MustCompile(dimensionName + r)
			subMatch := re.FindStringSubmatch(s)
			if len(subMatch) > 0 {
				l, _ := strconv.Atoi(numRegex.FindStringSubmatch(subMatch[0])[0])
				return Length(l) * lf.length
			}
		}
	}
	return 0
}
