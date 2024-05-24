package enigma

import (
	"testing"
)

func Test_runeToAlphabetIndex(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name      string
		selection rune
		want      int
	}{
		{
			name:      "A",
			selection: 'A',
			want:      0,
		},
		{
			name:      "B",
			selection: 'B',
			want:      1,
		},
		{
			name:      "Z",
			selection: 'Z',
			want:      25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeToAlphabetIndex(tt.selection); got != tt.want {
				t.Errorf("runeToAlphabetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alphabetIndexToRune(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name      string
		selection int
		want      rune
	}{
		{
			name:      "0",
			selection: 0,
			want:      'A',
		},
		{
			name:      "1",
			selection: 1,
			want:      'B',
		},
		{
			name:      "25",
			selection: 25,
			want:      'Z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabetIndexToRune(tt.selection); got != tt.want {
				t.Errorf("alphabetIndexToRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
