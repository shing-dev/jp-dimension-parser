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
			s: "幅11mm 奥25mm 高209mm",
			want: &Dimension{
				Width:  Length(11),
				Depth:  Length(25),
				Height: Length(209),
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
			s: "幅11~12cm",
			want: &Dimension{
				Width: Length(11) * Centimeter,
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
			s: "幅1cm",
			want: &Dimension{
				Width: Length(1) * Centimeter,
			},
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
