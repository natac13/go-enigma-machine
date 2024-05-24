package enigma

import "testing"

func TestNewReflector(t *testing.T) {
	// Test reflector B
	r, err := newReflector([]rune(REFLECTOR_B_WIRING))
	if r == nil {
		t.Error("NewReflector() returned nil")
	}
	if err != nil {
		t.Errorf("NewReflector() returned error: %v", err)
	}

	tests := []struct {
		input    rune
		expected rune
	}{
		{'A', 'Y'},
		{'B', 'R'},
		{'C', 'U'},
		{'D', 'H'},
		{'E', 'Q'},
		{'F', 'S'},
		{'G', 'L'},
		{'H', 'D'},
		{'I', 'P'},
		{'J', 'X'},
		{'K', 'N'},
		{'L', 'G'},
		{'M', 'O'},
		{'N', 'K'},
		{'O', 'M'},
		{'P', 'I'},
		{'Q', 'E'},
		{'R', 'B'},
		{'S', 'F'},
		{'T', 'Z'},
		{'U', 'C'},
		{'V', 'W'},
		{'W', 'V'},
		{'X', 'J'},
		{'Y', 'A'},
		{'Z', 'T'},
	}

	for _, test := range tests {
		if result, _ := r.transform(test.input); result != test.expected {
			t.Errorf("expected %c, got %c", test.expected, result)
		}
	}
}
