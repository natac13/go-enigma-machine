package enigma

import "testing"

func TestNewRotor(t *testing.T) {
	// Test rotor I
	r, err := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)
	if r == nil {
		t.Error("NewRotor() returned nil")
	}
	if err != nil {
		t.Errorf("NewRotor() returned error: %v", err)
	}

	if r.notch != 16 {
		t.Errorf("expected notch 16, got %d", r.notch)
	}

	if r.position != 0 {
		t.Errorf("expected position 0, got %d", r.position)
	}

	tests := []struct {
		input    rune
		expected rune
	}{
		{'A', 'E'},
		{'B', 'K'},
		{'C', 'M'},
		{'D', 'F'},
		{'E', 'L'},
		{'F', 'G'},
		{'G', 'D'},
		{'H', 'Q'},
		{'I', 'V'},
		{'J', 'Z'},
		{'K', 'N'},
		{'L', 'T'},
		{'M', 'O'},
		{'N', 'W'},
		{'O', 'Y'},
		{'P', 'H'},
		{'Q', 'X'},
		{'R', 'U'},
		{'S', 'S'},
		{'T', 'P'},
		{'U', 'A'},
		{'V', 'I'},
		{'W', 'B'},
		{'X', 'R'},
		{'Y', 'C'},
		{'Z', 'J'},
	}

	for _, test := range tests {
		if output, _ := r.transformForward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}

		if output, _ := r.transformBackward(test.expected); output != test.input {
			t.Errorf("expected %c, got %c", test.input, output)
		}
	}
}

func TestRotor_TransformForward(t *testing.T) {
	r, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)

	if r.position != 0 {
		t.Errorf("expected position 0, got %d", r.position)
	}

	tests := []struct {
		position string
		input    rune
		expected rune
	}{
		{"A", 'A', 'E'},
		{"B", 'A', 'J'},
		{"C", 'A', 'K'},
		{"D", 'A', 'C'},
		{"A", 'E', 'L'},
		{"B", 'E', 'F'},
		{"C", 'E', 'B'},
		{"B", 'K', 'S'},
	}

	for _, test := range tests {
		r.setPosition(test.position)
		if output, _ := r.transformForward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}
	}
}

func TestRotor_TransformBackward(t *testing.T) {
	r, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)

	if r.position != 0 {
		t.Errorf("expected position 0, got %d", r.position)
	}

	tests := []struct {
		position string
		input    rune
		expected rune
	}{
		{"A", 'E', 'A'},
		{"B", 'J', 'A'},
		{"C", 'K', 'A'},
		{"D", 'C', 'A'},
		{"A", 'L', 'E'},
		{"B", 'T', 'Q'},
	}

	for _, test := range tests {
		r.setPosition(test.position)
		if output, _ := r.transformBackward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}
	}
}

func TestNewRotor_InvalidWiring(t *testing.T) {
	_, err := NewRotor([]rune("ABC"), 'A')
	if err == nil {
		t.Error("NewRotor() did not return error for invalid wiring")
	}
}

func TestNewRotor_InvalidNotch(t *testing.T) {
	_, err := NewRotor([]rune(ROTOR_I_WIRING), '1')
	if err == nil {
		t.Error("NewRotor() did not return error for invalid notch")
	}
}

func TestRotor_Rotate(t *testing.T) {
	r, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)

	if r.position != 0 {
		t.Errorf("expected position 0, got %d", r.position)
	}

	rotateNext := r.rotate()
	if r.position != 1 {
		t.Errorf("expected position 1, got %d", r.position)
	}
	if rotateNext {
		t.Error("expected rotateNext to be false")
	}

	rotateNext = r.rotate()
	if r.position != 2 {
		t.Errorf("expected position 2, got %d", r.position)
	}
	if rotateNext {
		t.Error("expected rotateNext to be false")
	}

	r.setPosition("Z")
	r.rotate()
	if r.position != 0 {
		t.Errorf("expected position 0, got %d", r.position)
	}

	r.setPosition("Q")
	rotateNext = r.rotate()
	if r.position != 17 {
		t.Errorf("expected position 17, got %d", r.position)
	}
	if !rotateNext {
		t.Error("expected rotateNext to be true")
	}
}

func TestRotorsPathForward(t *testing.T) {
	rotor1, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)
	rotor2, _ := NewRotor([]rune(ROTOR_II_WIRING), ROTOR_II_NOTCH)
	rotor3, _ := NewRotor([]rune(ROTOR_III_WIRING), ROTOR_III_NOTCH)
	rotors := []*Rotor{
		rotor1,
		rotor2,
		rotor3,
	}

	tests := []struct {
		input              rune
		expectedFromRotor1 rune
		expectedFromRotor2 rune
		expectedFromRotor3 rune
	}{
		{'A', 'E', 'S', 'G'},
		{'B', 'K', 'L', 'V'},
	}

	for _, test := range tests {
		output, _ := rotors[0].transformForward(test.input)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}

		output, _ = rotors[1].transformForward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}

		output, _ = rotors[2].transformForward(output)
		if output != test.expectedFromRotor3 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor3, output)
		}
	}
}

func TestRotorsPathBackward(t *testing.T) {
	rotor1, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)
	rotor2, _ := NewRotor([]rune(ROTOR_II_WIRING), ROTOR_II_NOTCH)
	rotor3, _ := NewRotor([]rune(ROTOR_III_WIRING), ROTOR_III_NOTCH)
	rotors := []*Rotor{
		rotor1,
		rotor2,
		rotor3,
	}

	tests := []struct {
		input              rune
		expectedFromRotor3 rune
		expectedFromRotor2 rune
		expectedFromRotor1 rune
	}{
		{'G', 'S', 'E', 'A'},
		{'V', 'L', 'K', 'B'},
	}

	for _, test := range tests {
		output, _ := rotors[2].transformBackward(test.input)
		if output != test.expectedFromRotor3 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor3, output)
		}

		output, _ = rotors[1].transformBackward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}

		output, _ = rotors[0].transformBackward(output)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}
	}
}

func TestTransformForwardWithRotate(t *testing.T) {
	rotor1, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)
	rotor2, _ := NewRotor([]rune(ROTOR_II_WIRING), ROTOR_II_NOTCH)

	rotors := []*Rotor{
		rotor1,
		rotor2,
	}

	tests := []struct {
		input              rune
		expectedFromRotor1 rune
		expectedFromRotor2 rune
	}{
		{'A', 'J', 'B'},
		{'A', 'K', 'L'},
		{'A', 'C', 'D'},
	}

	for _, test := range tests {
		rotors[0].rotate()
		output, _ := rotors[0].transformForward(test.input)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}

		output, _ = rotors[1].transformForward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}
	}
}

func TestRotor_SetPosition(t *testing.T) {
	r, _ := NewRotor([]rune(ROTOR_I_WIRING), ROTOR_I_NOTCH)

	tests := []struct {
		letter   string
		position int
	}{
		{"A", 0},
		{"S", 18},
		{"Z", 25},
	}

	for _, test := range tests {
		if err := r.setPosition(test.letter); err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if r.position != test.position {
			t.Errorf("expected position %d, got %d", test.position, r.position)
		}
	}
}
