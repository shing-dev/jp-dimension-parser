package dimparser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s    string
		want *Dimension
	}{
		{
			s: "幅62cm×奥行73cm×高さ189cm",
			want: &Dimension{
				Width:  Length(62) * Centimeter,
				Depth:  Length(73) * Centimeter,
				Height: Length(189) * Centimeter,
			},
		},
		{
			s: "幅12.2cm×奥行1.5m×高さ18.9cm",
			want: &Dimension{
				Width:  Length(122),
				Depth:  Length(1500),
				Height: Length(189),
			},
		},
		{
			s: "幅 12cm×奥行　1.5 m×高さ: 18.9cm",
			want: &Dimension{
				Width:  Length(120),
				Depth:  Length(1500),
				Height: Length(189),
			},
		},
		{
			s: "幅11mm 奥行き25mm 高さ209〜250mm",
			want: &Dimension{
				Width:  Length(11),
				Depth:  Length(25),
				Height: Length(209),
			},
		},
		{
			s: "W11m×D24m×H99m",
			want: &Dimension{
				Width:  Length(11) * Meter,
				Depth:  Length(24) * Meter,
				Height: Length(99) * Meter,
			},
		},
		{
			s: "W９９mm×D８7cm×H6４m",
			want: &Dimension{
				Width:  Length(99) * Millimeter,
				Depth:  Length(87) * Centimeter,
				Height: Length(64) * Meter,
			},
		},
		{
			s: "10×11×12cm",
			want: &Dimension{
				Width:  Length(10) * Centimeter,
				Depth:  Length(11) * Centimeter,
				Height: Length(12) * Centimeter,
			},
		},
		{
			s: "幅10×奥行11×高さ12cm",
			want: &Dimension{
				Width:  Length(10) * Centimeter,
				Depth:  Length(11) * Centimeter,
				Height: Length(12) * Centimeter,
			},
		},
		{
			s: "W10×D11×H12 (mm)",
			want: &Dimension{
				Width:  Length(10) * Millimeter,
				Depth:  Length(11) * Millimeter,
				Height: Length(12) * Millimeter,
			},
		},
		{
			s: "幅11~12cm",
			want: &Dimension{
				Width: Length(11) * Centimeter,
			},
		},
		{
			s: "幅11-12cm",
			want: &Dimension{
				Width: Length(11) * Centimeter,
			},
		},
		{
			s: "幅1cm",
			want: &Dimension{
				Width: Length(1) * Centimeter,
			},
		},
		{
			s:    "幅62×奥行73×高さ189",
			want: nil,
		},
		{
			s:    "幅62dm×奥行73am×高さ189pm",
			want: nil,
		},
		{
			s:    "床面までの高さ10cm",
			want: nil,
		},
		{
			s:    "座面高さ24cm",
			want: nil,
		},
		{
			s:    "寸法に関する文字なし",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := Parse(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
