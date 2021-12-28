package jp_dimension_parser

import "testing"

func Test_analyze(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normalize full-width character",
			s:    "１２３",
			want: "123",
		},
		{
			name: "replace misleading words",
			s:    "座面高さ123cm、高さ20cm",
			want: "#123cm、高さ20cm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyze(tt.s); got != tt.want {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
