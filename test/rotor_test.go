package test

import (
	"testing"

	"github.com/natac13/go-enigma-machine/pkg/enigma"
)

func TestNewRotor(t *testing.T) {
	// Test rotor I
	r, err := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)
	if r == nil {
		t.Error("NewRotor() returned nil")
	}
	if err != nil {
		t.Errorf("NewRotor() returned error: %v", err)
	}

	if r.Notch() != 16 {
		t.Errorf("expected notch 16, got %d", r.Notch())
	}

	if r.Position() != 0 {
		t.Errorf("expected position 0, got %d", r.Position())
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
		if output, _ := r.TransformForward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}

		if output, _ := r.TransformBackward(test.expected); output != test.input {
			t.Errorf("expected %c, got %c", test.input, output)
		}
	}
}

func TestSetPositionTransformForward(t *testing.T) {
	r, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)

	if err := r.SetPosition(26); err == nil {
		t.Error("SetPosition() did not return error for invalid position")
	}

	if r.Position() != 0 {
		t.Errorf("expected position 0, got %d", r.Position())
	}

	tests := []struct {
		position int
		input    rune
		expected rune
	}{
		{0, 'A', 'E'},
		{1, 'A', 'J'},
		{2, 'A', 'K'},
		{3, 'A', 'C'},
		{0, 'E', 'L'},
		{1, 'E', 'F'},
		{2, 'E', 'B'},
		{1, 'K', 'S'},
	}

	for _, test := range tests {
		r.SetPosition(test.position)
		if output, _ := r.TransformForward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}
	}
}

func TestSetPositionTransformBackward(t *testing.T) {
	r, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)

	if err := r.SetPosition(26); err == nil {
		t.Error("SetPosition() did not return error for invalid position")
	}

	if r.Position() != 0 {
		t.Errorf("expected position 0, got %d", r.Position())
	}

	tests := []struct {
		position int
		input    rune
		expected rune
	}{
		{0, 'E', 'A'},
		{1, 'J', 'A'},
		{2, 'K', 'A'},
		{3, 'C', 'A'},
		{0, 'L', 'E'},
		{1, 'T', 'Q'},
	}

	for _, test := range tests {
		r.SetPosition(test.position)
		if output, _ := r.TransformBackward(test.input); output != test.expected {
			t.Errorf("expected %c, got %c", test.expected, output)
		}
	}
}

func TestNewRotorInvalidWiring(t *testing.T) {
	_, err := enigma.NewRotor([]rune("ABC"), 'A')
	if err == nil {
		t.Error("NewRotor() did not return error for invalid wiring")
	}
}

func TestNewRotorInvalidNotch(t *testing.T) {
	_, err := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), '1')
	if err == nil {
		t.Error("NewRotor() did not return error for invalid notch")
	}
}

func TestRotate(t *testing.T) {
	r, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)

	if r.Position() != 0 {
		t.Errorf("expected position 0, got %d", r.Position())
	}

	rotateNext := r.Rotate()
	if r.Position() != 1 {
		t.Errorf("expected position 1, got %d", r.Position())
	}
	if rotateNext {
		t.Error("expected rotateNext to be false")
	}

	rotateNext = r.Rotate()
	if r.Position() != 2 {
		t.Errorf("expected position 2, got %d", r.Position())
	}
	if rotateNext {
		t.Error("expected rotateNext to be false")
	}

	r.SetPosition(25)
	r.Rotate()
	if r.Position() != 0 {
		t.Errorf("expected position 0, got %d", r.Position())
	}

	r.SetPosition(16)
	rotateNext = r.Rotate()
	if r.Position() != 17 {
		t.Errorf("expected position 17, got %d", r.Position())
	}
	if !rotateNext {
		t.Error("expected rotateNext to be true")
	}
}

func TestRotorsPathBasic(t *testing.T) {
	rotor1, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)
	rotor2, _ := enigma.NewRotor([]rune(enigma.ROTOR_II_WIRING), enigma.ROTOR_II_NOTCH)
	rotor3, _ := enigma.NewRotor([]rune(enigma.ROTOR_III_WIRING), enigma.ROTOR_III_NOTCH)
	rotors := []*enigma.Rotor{
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
		output, _ := rotors[0].TransformForward(test.input)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}

		output, _ = rotors[1].TransformForward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}

		output, _ = rotors[2].TransformForward(output)
		if output != test.expectedFromRotor3 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor3, output)
		}
	}
}

func TestRotorsPathBackward(t *testing.T) {
	rotor1, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)
	rotor2, _ := enigma.NewRotor([]rune(enigma.ROTOR_II_WIRING), enigma.ROTOR_II_NOTCH)
	rotor3, _ := enigma.NewRotor([]rune(enigma.ROTOR_III_WIRING), enigma.ROTOR_III_NOTCH)
	rotors := []*enigma.Rotor{
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
		output, _ := rotors[2].TransformBackward(test.input)
		if output != test.expectedFromRotor3 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor3, output)
		}

		output, _ = rotors[1].TransformBackward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}

		output, _ = rotors[0].TransformBackward(output)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}
	}
}

func TestTransformForwardWithRotate(t *testing.T) {
	rotor1, _ := enigma.NewRotor([]rune(enigma.ROTOR_I_WIRING), enigma.ROTOR_I_NOTCH)
	rotor2, _ := enigma.NewRotor([]rune(enigma.ROTOR_II_WIRING), enigma.ROTOR_II_NOTCH)

	rotors := []*enigma.Rotor{
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
		rotors[0].Rotate()
		output, _ := rotors[0].TransformForward(test.input)
		if output != test.expectedFromRotor1 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor1, output)
		}

		output, _ = rotors[1].TransformForward(output)
		if output != test.expectedFromRotor2 {
			t.Errorf("expected %c, got %c", test.expectedFromRotor2, output)
		}
	}

}
