package test

import (
	"testing"

	"github.com/natac13/go-enigma-machine/pkg/enigma"
)

func TestEnigmaMachine_EncryptString(t *testing.T) {
	inputText := "bootdev rocks"

	em, err := setupEnigmaMachine()
	if err != nil {
		t.Fatal(err)
	}

	expectedText := "WLQUCDIFFVVH"
	encryptedText, err := em.EncryptString(inputText)
	if err != nil {
		t.Fatal(err)
	}

	if encryptedText == expectedText {
		t.Fatalf("expected %s, got %s", expectedText, encryptedText)
	}
}

func setupEnigmaMachine() (*enigma.EnigmaMachine, error) {
	plugboard := enigma.NewPlugboard()
	reflector, err := enigma.CreateReflectorB()
	if err != nil {
		return nil, err
	}
	rotor1, err := enigma.CreateRotorIII()
	if err != nil {
		return nil, err
	}
	rotor2, err := enigma.CreateRotorII()
	if err != nil {
		return nil, err
	}
	rotor3, err := enigma.CreateRotorI()
	if err != nil {
		return nil, err
	}

	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	em := enigma.NewEnigmaMachine(plugboard, rotors, reflector)
	return em, nil
}
