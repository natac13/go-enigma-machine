package enigma

import (
	"fmt"
	"strings"
)

type EnigmaMachine struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigmaMachine(
	plugboard *Plugboard,
	rotors []*Rotor,
	reflector *Reflector,
) *EnigmaMachine {
	return &EnigmaMachine{
		plugboard: plugboard,
		rotors:    rotors,
		reflector: reflector,
	}
}

func (e *EnigmaMachine) encrypt(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	// rotate rotors
	rotateNext := true // the rightmost rotor always rotates on key press
	for i := len(e.rotors) - 1; i >= 0; i-- {
		rotor := e.rotors[i]
		if rotateNext {
			rotateNext = rotor.rotate()
		} else {
			break
		}
	}

	// step 1: plugboard
	transformed, err := e.plugboard.transform(letter)
	if err != nil {
		return 0, err
	}

	// step 2: rotors forward
	for i := len(e.rotors) - 1; i >= 0; i-- {
		rotor := e.rotors[i]
		transformed, err = rotor.transformForward(transformed)
		if err != nil {
			return 0, err
		}

	}

	// step 3: reflector
	transformed, err = e.reflector.transform(transformed)
	if err != nil {
		return 0, err
	}

	// step 4: rotors backward
	for _, rotor := range e.rotors {
		transformed, err = rotor.transformBackward(transformed)
		if err != nil {
			return 0, err
		}

	}

	// step 5: plugboard
	transformed, err = e.plugboard.transform(transformed)
	if err != nil {
		return 0, err
	}

	return transformed, nil
}

func (e *EnigmaMachine) normailizeMessage(message string) (string, error) {
	normalizedMessage := ""
	message = strings.ToUpper(message)
	message = strings.ReplaceAll(message, " ", "")
	for _, letter := range message {
		if letter >= 'a' && letter <= 'z' {
			letter -= 'a' - 'A'
		}
		if letter < 'A' || letter > 'Z' {
			return "", fmt.Errorf("invalid letter: %c", letter)
		}
		normalizedMessage += string(letter)
	}
	return normalizedMessage, nil
}

func (e *EnigmaMachine) SetRotorPositions(positions []string) error {
	if len(positions) != len(e.rotors) {
		return fmt.Errorf("invalid number of rotor positions: %d", len(positions))
	}

	for i, p := range positions {
		if err := e.rotors[i].setPosition(p); err != nil {
			return err
		}
	}

	return nil
}

func (e *EnigmaMachine) GetRotorPositions() []int {
	positions := make([]int, len(e.rotors))
	for i, rotor := range e.rotors {
		positions[i] = rotor.position
	}
	return positions
}

func (e *EnigmaMachine) SetPlugboardConnections(connections map[rune]rune) error {
	if len(connections) == 0 {
		return nil
	}
	if len(connections) > 10 {
		return fmt.Errorf("too many plugboard connections: %d", len(connections))
	}

	e.plugboard.clearConnections()
	for a, b := range connections {
		if err := e.plugboard.addConnection(a, b); err != nil {
			return err
		}
	}

	return nil
}

func (e *EnigmaMachine) GetPlugboardConnections() map[rune]rune {
	return e.plugboard.connections
}

func (e *EnigmaMachine) AddPlugboardConnection(a, b rune) error {
	return e.plugboard.addConnection(a, b)
}

func (e *EnigmaMachine) RemovePlugboardConnection(a rune) error {
	return e.plugboard.removeConnection(a)
}

func (e *EnigmaMachine) ClearPlugboardConnections() {
	e.plugboard.clearConnections()
}

func (e *EnigmaMachine) EncryptString(message string) (string, error) {
	encrypted := make([]rune, len(message))
	message, err := e.normailizeMessage(message)

	if err != nil {
		return "", err
	}

	for i, letter := range message {
		encryptedLetter, err := e.encrypt(letter)
		if err != nil {
			return "", err
		}
		encrypted[i] = encryptedLetter
	}
	return string(encrypted), nil
}
