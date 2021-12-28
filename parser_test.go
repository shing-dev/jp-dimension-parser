package jp_dimension_parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		s     string
		want  *Dimension
		want1 bool
	}{
		{
			s: "幅62cm×奥行73cm×高さ189cm",
			want: &Dimension{
				Width:  Length(62) * Centimeter,
				Depth:  Length(73) * Centimeter,
				Height: Length(189) * Centimeter,
			},
			want1: true,
		},
		{
			s: "幅11mm 奥行き25mm 高209mm",
			want: &Dimension{
				Width:  Length(11),
				Depth:  Length(25),
				Height: Length(209),
			},
			want1: true,
		},
		{
			s: "幅11mm 奥行き25mm 高さ209〜250mm",
			want: &Dimension{
				Width:  Length(11),
				Depth:  Length(25),
				Height: Length(209),
			},
			want1: true,
		},
		{
			s: "幅11~12cm",
			want: &Dimension{
				Width: Length(11) * Centimeter,
			},
			want1: true,
		},
		{
			s: "W11m×D24m×H99m",
			want: &Dimension{
				Width:  Length(11) * Meter,
				Depth:  Length(24) * Meter,
				Height: Length(99) * Meter,
			},
			want1: true,
		},
		{
			s: "W９９mm×D８7cm×H6４m",
			want: &Dimension{
				Width:  Length(99) * Millimeter,
				Depth:  Length(87) * Centimeter,
				Height: Length(64) * Meter,
			},
			want1: true,
		},
		{
			s: "幅1cm",
			want: &Dimension{
				Width: Length(1) * Centimeter,
			},
			want1: true,
		},
		{
			s:     "座面高さ24cm",
			want:  nil,
			want1: false,
		},
		{
			s:     "寸法に関する文字なし",
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
