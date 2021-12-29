package dimparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	lengthFormats = []struct {
		format string
		length Length
	}{
		{format: "mm", length: Millimeter},
		{format: "cm", length: Centimeter},
		{format: "m", length: Meter},
	}

	widthFormats  = []string{"幅", "width", "w", "W"}
	depthFormats  = []string{"奥行き", "奥行", "奥", "長さ", "depth", "D", "d"}
	heightFormats = []string{"高さ", "高", "height", "H", "h"}
)

var floatNumRegex = regexp.MustCompile(`([0-9]*[.])?[0-9]+`)

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

	if dim := parseAllDimensions(s); dim != nil {
		return dim
	}
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
	if length := parseDimension(widthFormats, s); length > 0 {
		return length
	}
	return 0
}

func parseDepth(s string) Length {
	if length := parseDimension(depthFormats, s); length > 0 {
		return length
	}
	return 0
}

func parseHeight(s string) Length {
	length := parseDimension(heightFormats, s)
	if length > 0 {
		return length
	}
	return 0
}

func parseAllDimensions(s string) *Dimension {
	widthNameRegexOr := dimensionNamesToRegexOr(widthFormats)
	depthNameRegexOr := dimensionNamesToRegexOr(depthFormats)
	heightNameRegexOr := dimensionNamesToRegexOr(heightFormats)
	for _, lf := range lengthFormats {
		regexString := fmt.Sprintf(
			`%s?([0-9]*[.])?[0-9]+([ ]*)[×xX]%s?([0-9]*[.])?[0-9]+([ ]*)[×X]%s?([0-9]*[.])?[0-9]+([ ]*)?[\(]?%s[\)]?`,
			widthNameRegexOr,
			depthNameRegexOr,
			heightNameRegexOr,
			lf.format,
		)
		re := regexp.MustCompile(regexString)
		subMatch := re.FindStringSubmatch(s)
		if len(subMatch) > 0 {
			lengthStrings := floatNumRegex.FindAllStringSubmatch(subMatch[0], 3)
			w, _ := strconv.ParseFloat(lengthStrings[0][0], 64)
			d, _ := strconv.ParseFloat(lengthStrings[1][0], 64)
			h, _ := strconv.ParseFloat(lengthStrings[2][0], 64)
			return &Dimension{
				Width:  Length(w * float64(lf.length)),
				Depth:  Length(d * float64(lf.length)),
				Height: Length(h * float64(lf.length)),
			}
		}
	}

	return nil
}

func parseDimension(dimensionNames []string, s string) Length {
	regexes := buildLengthRegexes(dimensionNames)
	for _, r := range regexes {
		re := regexp.MustCompile(r.regexString)
		subMatch := re.FindStringSubmatch(s)
		if len(subMatch) > 0 {
			l, _ := strconv.ParseFloat(floatNumRegex.FindStringSubmatch(subMatch[0])[0], 64)
			return Length(l * float64(r.length))
		}
	}
	return 0
}

type lengthRegex struct {
	regexString string
	length      Length
}

func buildLengthRegexes(dimensionNames []string) []*lengthRegex {
	var regexes []*lengthRegex

	dimNameRegexOr := dimensionNamesToRegexOr(dimensionNames)
	for _, lf := range lengthFormats {
		regexStrings := []string{
			fmt.Sprintf(`%s([ :]*)?([0-9]*[.])?[0-9]+([ ]*)?%s`, dimNameRegexOr, lf.format),
			fmt.Sprintf(`%s([ :]*)?([0-9]*[.])?[0-9]+[〜~-][0-9]+([ ]*)?%s`, dimNameRegexOr, lf.format),
		}
		for _, r := range regexStrings {
			regexes = append(regexes, &lengthRegex{
				regexString: r,
				length:      lf.length,
			})
		}
	}

	return regexes
}

// dimensionNamesToRegexOr will convert dimension names to OR regex string for the dimension
// e.g. []string{"幅", "W"} to (幅 | W)
func dimensionNamesToRegexOr(dimensionNames []string) string {
	regexOr := strings.Join(dimensionNames, ")|(")
	regexOr = fmt.Sprintf("((%s))", regexOr)
	return regexOr
}
