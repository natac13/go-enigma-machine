package test

import (
	"testing"

	"github.com/natac13/go-enigma-machine/pkg/enigma"
)

func TestNewPlugboard(t *testing.T) {
	p := enigma.NewPlugboard()
	if p == nil {
		t.Error("NewPlugboard() returned nil")
	}

	if p.CountConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.CountConnections())
	}

	tests := []struct {
		input    rune
		expected rune
	}{
		{'A', 'A'},
		{'B', 'B'},
		{'C', 'C'},
		{'D', 'D'},
		{'E', 'E'},
		{'F', 'F'},
		{'G', 'G'},
		{'H', 'H'},
		{'I', 'I'},
		{'J', 'J'},
		{'K', 'K'},
		{'L', 'L'},
		{'M', 'M'},
		{'N', 'N'},
		{'O', 'O'},
		{'P', 'P'},
		{'Q', 'Q'},
		{'R', 'R'},
		{'S', 'S'},
		{'T', 'T'},
		{'U', 'U'},
		{'V', 'V'},
		{'W', 'W'},
		{'X', 'X'},
		{'Y', 'Y'},
		{'Z', 'Z'},
	}

	for _, test := range tests {
		if p.Transform(test.input) != test.expected {
			t.Errorf("expected %c, got %c", test.expected, p.Transform(test.input))
		}
	}
}

func TestAddConnection(t *testing.T) {
	p := enigma.NewPlugboard()

	tests := []struct {
		a rune
		b rune
	}{
		{'A', 'B'},
		{'C', 'D'},
		{'E', 'F'},
		{'G', 'H'},
		{'I', 'J'},
		{'K', 'L'},
		{'M', 'N'},
		{'O', 'P'},
		{'Q', 'R'},
		{'S', 'T'},
	}

	for _, test := range tests {
		if err := p.AddConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if p.Transform(test.a) != test.b {
			t.Errorf("expected %c, got %c", test.b, p.Transform(test.a))
		}

		if p.Transform(test.b) != test.a {
			t.Errorf("expected %c, got %c", test.a, p.Transform(test.b))
		}
	}
}
