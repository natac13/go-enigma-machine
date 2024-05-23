package enigma

import "fmt"

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

func (e *EnigmaMachine) Encrypt(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	// rotate rotors
	rotateNext := true // the rightmost rotor always rotates on key press
	for i := len(e.rotors) - 1; i >= 0; i-- {
		rotor := e.rotors[i]
		if rotateNext {
			rotateNext = rotor.Rotate()
		} else {
			break
		}
	}

	// step 1: plugboard
	transformed, err := e.plugboard.Transform(letter)
	if err != nil {
		return 0, err
	}
	fmt.Println("plugboard encryption: ", string(transformed))

	// step 2: rotors forward
	for i := len(e.rotors) - 1; i >= 0; i-- {
		rotor := e.rotors[i]
		transformed, err = rotor.TransformForward(transformed)
		if err != nil {
			return 0, err
		}
		fmt.Println("rotor ", i, ": ", string(transformed))
	}

	// step 3: reflector
	transformed, err = e.reflector.Transform(transformed)
	if err != nil {
		return 0, err
	}
	fmt.Println("reflector: ", string(transformed))

	// step 4: rotors backward
	for i, rotor := range e.rotors {
		transformed, err = rotor.TransformBackward(transformed)
		if err != nil {
			return 0, err
		}
		fmt.Println("rotor ", i, ": ", string(transformed))
	}

	// step 5: plugboard
	transformed, err = e.plugboard.Transform(transformed)
	if err != nil {
		return 0, err
	}
	fmt.Println("plugboard 2: ", string(transformed))

	return transformed, nil
}

func (e *EnigmaMachine) SetRotorPositions(positions []int) error {
	if len(positions) != len(e.rotors) {
		return fmt.Errorf("invalid number of rotor positions: %d", len(positions))
	}

	for i, position := range positions {
		if err := e.rotors[i].SetPosition(position); err != nil {
			return err
		}
	}

	return nil
}

func (e *EnigmaMachine) GetRotorPositions() []int {
	positions := make([]int, len(e.rotors))
	for i, rotor := range e.rotors {
		positions[i] = rotor.Position()
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

	e.plugboard.ClearConnections()
	for a, b := range connections {
		if err := e.plugboard.AddConnection(a, b); err != nil {
			return err
		}
	}

	return nil
}

func (e *EnigmaMachine) GetPlugboardConnections() map[rune]rune {
	return e.plugboard.GetConnections()
}

func (e *EnigmaMachine) AddPlugboardConnection(a, b rune) error {
	return e.plugboard.AddConnection(a, b)
}

func (e *EnigmaMachine) RemovePlugboardConnection(a rune) error {
	return e.plugboard.RemoveConnection(a)
}

func (e *EnigmaMachine) ClearPlugboardConnections() {
	e.plugboard.ClearConnections()
}
