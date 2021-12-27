package jp_dimension_parser

type Length int64

const (
	Millimeter Length = 1
	Centimeter        = Millimeter * 10
	Meter             = Centimeter * 100
	KiloMeter         = Meter * 1000
)

func (l Length) Millimeter() int64 {
	return int64(l)
}

func (l Length) Centimeter() float64 {
	cm := l / Centimeter
	mm := l % Centimeter
	return float64(cm) + float64(mm)/10
}

func (l Length) Meter() float64 {
	cm := l / Meter
	mm := l % Meter
	return float64(cm) + float64(mm)/1e3
}

func (l Length) KiloMeter() float64 {
	cm := l / KiloMeter
	mm := l % KiloMeter
	return float64(cm) + float64(mm)/1e6
}

