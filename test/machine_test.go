package test

import (
	"errors"
	"strings"
	"testing"

	"github.com/natac13/go-enigma-machine/pkg/enigma"
)

func TestEnigmaMachineEncrypt(t *testing.T) {
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
		encrypted, err := em.Encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.SetRotorPositions([]int{0, 0, 0})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'F', 'A'},
		{'T', 'A'},
		{'Z', 'A'},
	}

	for _, test := range tests {
		encrypted, err := em.Encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}
}

func TestEnigmaMachineEncrypt_WithPlugboard(t *testing.T) {
	plugboard := enigma.NewPlugboard()
	reflector, err := enigma.NewReflector([]rune(enigma.REFLECTOR_B_WIRING))
	if err != nil {
		t.Fatal(err)
	}
	rotor1, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_III_WIRING),
		enigma.ROTOR_III_NOTCH,
	)
	if err != nil {
		t.Fatal(err)
	}
	rotor2, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_II_WIRING),
		enigma.ROTOR_II_NOTCH,
	)
	if err != nil {
		t.Fatal(err)
	}
	rotor3, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_I_WIRING),
		enigma.ROTOR_I_NOTCH,
	)
	if err != nil {
		t.Fatal(err)
	}

	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	em := enigma.NewEnigmaMachine(plugboard, rotors, reflector)

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
		encrypted, err := em.Encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.SetRotorPositions([]int{0, 0, 0})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'M', 'A'},
		{'T', 'A'},
		{'Z', 'A'},
	}

	for _, test := range tests {
		encrypted, err := em.Encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}

	em.AddPlugboardConnection('N', 'Z')
	em.SetRotorPositions([]int{0, 0, 0})

	tests = []struct {
		input    rune
		expected rune
	}{
		{'A', 'M'},
		{'A', 'T'},
		{'A', 'N'},
	}

	for _, test := range tests {
		encrypted, err := em.Encrypt(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if encrypted != test.expected {
			t.Fatalf("expected %c, got %c", test.expected, encrypted)
		}
	}
}

func TestEnigmaMachine_WithText(t *testing.T) {
	inputText := "bootdev rocks"

	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	expectedText := "wlqucdiffvvh"
	encryptedText := ""

	normailzedInputText := strings.ReplaceAll(strings.ToUpper(inputText), " ", "")

	for _, letter := range normailzedInputText {
		r := rune(letter)
		if r == ' ' {
			encryptedText += " "
			continue
		}

		encrypted, err := em.Encrypt(r)
		if err != nil {
			t.Fatal(err)
		}

		encryptedText += string(encrypted)
	}

	if encryptedText == expectedText {
		t.Fatalf("expected %s, got %s", expectedText, encryptedText)
	}

}

func TestNormalizeAndValidateIncomingMessage(t *testing.T) {
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
		normalized, err := em.NormailizeAndValidateIncomingMessage(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if normalized != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, normalized)
		}
	}
}

func TestNormailzeAndValidateIncomingMessage_Invalid(t *testing.T) {
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
		_, err := em.NormailizeAndValidateIncomingMessage(test.input)
		if err == nil {
			t.Fatal(errors.New("expected error, got nil"))
		}

		if err.Error() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, err.Error())
		}
	}
}

func setupEnigmaMachine() (*enigma.EnigmaMachine, error) {
	plugboard := enigma.NewPlugboard()
	reflector, err := enigma.NewReflector([]rune(enigma.REFLECTOR_B_WIRING))
	if err != nil {
		return nil, err
	}
	rotor1, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_III_WIRING),
		enigma.ROTOR_III_NOTCH,
	)
	if err != nil {
		return nil, err
	}
	rotor2, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_II_WIRING),
		enigma.ROTOR_II_NOTCH,
	)
	if err != nil {
		return nil, err
	}
	rotor3, err := enigma.NewRotor(
		[]rune(enigma.ROTOR_I_WIRING),
		enigma.ROTOR_I_NOTCH,
	)
	if err != nil {
		return nil, err
	}

	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	em := enigma.NewEnigmaMachine(plugboard, rotors, reflector)
	return em, nil
}
