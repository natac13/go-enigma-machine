package enigma

import (
	"errors"
	"testing"
)

func TestEnigmaMachine_Encrypt(t *testing.T) {
	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    rune
		expected rune
	}{
		{'A', 'F'},
		{'A', 'T'},
		{'A', 'Z'},
	}

	for _, test := range tests {
		encrypted, err := em.encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.SetRotorPositions([]string{"A", "A", "A"})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'F', 'A'},
		{'T', 'A'},
		{'Z', 'A'},
	}

	for _, test := range tests {
		encrypted, err := em.encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}
}

func TestEnigmaMachine_Encrypt_WithPlugboard(t *testing.T) {
	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	em.SetPlugboardConnections(map[rune]rune{
		'M': 'F',
	})

	tests := []struct {
		input    rune
		expected rune
	}{
		{'A', 'M'},
		{'A', 'T'},
		{'A', 'Z'},
	}

	for _, test := range tests {
		encrypted, err := em.encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.SetRotorPositions([]string{"A", "A", "A"})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'M', 'A'},
		{'T', 'A'},
		{'Z', 'A'},
	}

	for _, test := range tests {
		encrypted, err := em.encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.AddPlugboardConnection('N', 'Z')
	em.SetRotorPositions([]string{"A", "A", "A"})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'A', 'M'},
		{'A', 'T'},
		{'A', 'N'},
	}

	for _, test := range tests {
		encrypted, err := em.encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}
}

func TestEnigmaMachine_NormalizeMessage(t *testing.T) {
	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"bootdev rocks", "BOOTDEVROCKS"},
		{"BootDev Rocks", "BOOTDEVROCKS"},
	}

	for _, test := range tests {
		normalized, err := em.normailizeMessage(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if normalized != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, normalized)
		}
	}
}

func TestEnigmaMachine_NormailzeMessage_Invalid(t *testing.T) {
	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"BootDev Rocks!", "invalid letter: !"},
		{"BootDev Rocks@#", "invalid letter: @"},
		{"BootDev Rocks#123", "invalid letter: #"},
		{"BootDev Rocks123", "invalid letter: 1"},
	}

	for _, test := range tests {
		_, err := em.normailizeMessage(test.input)
		if err == nil {
			t.Fatal(errors.New("expected error, got nil"))
		}

		if err.Error() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, err.Error())
		}
	}
}

// for testing purposes only
func setupEnigmaMachine() (*EnigmaMachine, error) {
	plugboard := NewPlugboard()
	reflector, err := CreateReflectorB()
	if err != nil {
		return nil, err
	}
	rotor1, err := CreateRotorIII()
	if err != nil {
		return nil, err
	}
	rotor2, err := CreateRotorII()
	if err != nil {
		return nil, err
	}
	rotor3, err := CreateRotorI()
	if err != nil {
		return nil, err
	}

	rotors := []*Rotor{rotor1, rotor2, rotor3}

	em := NewEnigmaMachine(plugboard, rotors, reflector)
	return em, nil
}
