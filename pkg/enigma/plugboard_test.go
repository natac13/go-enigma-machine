package enigma

import "testing"

func TestNewPlugboard(t *testing.T) {
	p := NewPlugboard()
	if p == nil {
		t.Error("NewPlugboard() returned nil")
	}

	if p.countConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.countConnections())
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
		if val, _ := p.transform(test.input); val != test.expected {
			t.Errorf("expected %c, got %c", test.expected, val)
		}
	}
}

func TestAddConnection(t *testing.T) {
	p := NewPlugboard()

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
		if err := p.addConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if val, _ := p.transform(test.a); val != test.b {
			t.Errorf("expected %c, got %c", test.b, val)
		}

		if val, _ := p.transform(test.b); val != test.a {
			t.Errorf("expected %c, got %c", test.a, val)
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

	p := NewPlugboard()

	for _, test := range tests {
		if err := p.addConnection(test.a, test.b); err != nil {
			if err.Error() != test.error {
				t.Errorf("expected %q, got %q", test.error, err.Error())
			}
		}
	}
}

func TestRemoveConnection(t *testing.T) {
	p := NewPlugboard()

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
		if err := p.addConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.countConnections() != 4 {
		t.Errorf("expected 8 connections, got %d", p.countConnections())
	}

	for _, test := range tests {
		if err := p.removeConnection(test.a); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if val, _ := p.transform(test.a); val != test.a {
			t.Errorf("expected %c, got %c", test.a, val)
		}

		if val, _ := p.transform(test.b); val != test.b {
			t.Errorf("expected %c, got %c", test.b, val)
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

	p := NewPlugboard()

	for _, test := range tests {
		if err := p.removeConnection(test.a); err != nil {
			if err.Error() != test.error {
				t.Errorf("expected %q, got %q", test.error, err.Error())
			}
		}
	}
}

func TestClearConnections(t *testing.T) {
	p := NewPlugboard()

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
		if err := p.addConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.countConnections() != 4 {
		t.Errorf("expected 4 connections, got %d", p.countConnections())
	}

	p.clearConnections()

	if p.countConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.countConnections())
	}
}

func TestCountConnections(t *testing.T) {
	p := NewPlugboard()

	if p.countConnections() != 0 {
		t.Errorf("expected 0 connections, got %d", p.countConnections())
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
		if err := p.addConnection(test.a, test.b); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if p.countConnections() != 4 {
		t.Errorf("expected 4 connections, got %d", p.countConnections())
	}
}
