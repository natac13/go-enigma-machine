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

func TestAddConnectionErrors(t *testing.T) {
	tests := []struct {
		a     rune
		b     rune
		error string
	}{
		{'A', 'A', "cannot connect a letter to itself: A A"},
		{'A', 'B', ""},
		{'A', 'C', "letter A is already connected"},
		{'B', 'C', "letter B is already connected"},
		{'a', 'B', "invalid connection: a B"},
		{'A', 'b', "invalid connection: A b"},
	}

	p := enigma.NewPlugboard()

	for _, test := range tests {
		if err := p.AddConnection(test.a, test.b); err != nil {
			if err.Error() != test.error {
				t.Errorf("expected %q, got %q", test.error, err.Error())
			}
		}
	}
}

func TestRemoveConnection(t *testing.T) {
	p := enigma.NewPlugboard()

	tests := []struct {
		a rune
		b rune
	}{
		{'A', 'B'},
		{'C', 'D'},
		{'E', 'F'},
		{'G', 'H'},
	}

	for _, test := range tests {
		if err := p.AddConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.CountConnections() != 4 {
		t.Errorf("expected 8 connections, got %d", p.CountConnections())
	}

	for _, test := range tests {
		if err := p.RemoveConnection(test.a); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if p.Transform(test.a) != test.a {
			t.Errorf("expected %c, got %c", test.a, p.Transform(test.a))
		}

		if p.Transform(test.b) != test.b {
			t.Errorf("expected %c, got %c", test.b, p.Transform(test.b))
		}
	}
}

func TestRemoveConnectionErrors(t *testing.T) {
	tests := []struct {
		a     rune
		error string
	}{
		{'A', "letter A is not connected"},
		{'a', "invalid connection: a"},
	}

	p := enigma.NewPlugboard()

	for _, test := range tests {
		if err := p.RemoveConnection(test.a); err != nil {
			if err.Error() != test.error {
				t.Errorf("expected %q, got %q", test.error, err.Error())
			}
		}
	}
}

func TestClearConnections(t *testing.T) {
	p := enigma.NewPlugboard()

	tests := []struct {
		a rune
		b rune
	}{
		{'A', 'B'},
		{'C', 'D'},
		{'E', 'F'},
		{'G', 'H'},
	}

	for _, test := range tests {
		if err := p.AddConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.CountConnections() != 4 {
		t.Errorf("expected 4 connections, got %d", p.CountConnections())
	}

	p.ClearConnections()

	if p.CountConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.CountConnections())
	}
}

func TestCountConnections(t *testing.T) {
	p := enigma.NewPlugboard()

	if p.CountConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.CountConnections())
	}

	tests := []struct {
		a rune
		b rune
	}{
		{'A', 'B'},
		{'C', 'D'},
		{'E', 'F'},
		{'G', 'H'},
	}

	for _, test := range tests {
		if err := p.AddConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.CountConnections() != 4 {
		t.Errorf("expected 4 connections, got %d", p.CountConnections())
	}
}
