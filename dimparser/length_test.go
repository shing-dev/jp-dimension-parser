package dimparser

import "testing"

func TestLength_Millimeters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    Length
		want int64
	}{
		{
			name: "returns 1 for 1mm",
			l:    1,
			want: 1,
		},
		{
			name: "returns 10 for 1cm",
			l:    Centimeter,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Millimeters(); got != tt.want {
				t.Errorf("Millimeters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength_Centimeters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    Length
		want float64
	}{
		{
			name: "returns 0.1 for 1mm",
			l:    Millimeter,
			want: 0.1,
		},
		{
			name: "returns 1 for 1cm",
			l:    Centimeter,
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Centimeters(); got != tt.want {
				t.Errorf("Centimeters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength_Meters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    Length
		want float64
	}{
		{
			name: "return 0.001 for 1mm",
			l:    Millimeter,
			want: 0.001,
		},
		{
			name: "returns 0.01 for 1cm",
			l:    Centimeter,
			want: 0.01,
		},
		{
			name: "returns 1 for 1m",
			l:    Meter,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Meters(); got != tt.want {
				t.Errorf("Meters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength_KiloMeters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    Length
		want float64
	}{
		{
			name: "return 0.001 for 1m",
			l:    Meter,
			want: 0.001,
		},
		{
			name: "returns 1 for 1m",
			l:    KiloMeter,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Kilometers(); got != tt.want {
				t.Errorf("Kilometers() = %v, want %v", got, tt.want)
			}
		})
	}
}
