package enigma

import (
	"fmt"
	"slices"
)

type Rotor struct {
	wiring   []rune
	notch    int
	position int
}

func runeToAlphabetIndex(r rune) int {
	return int(r - 'A')
}

func alphabetIndexToRune(i int) rune {
	return rune(i + 'A')
}

func NewRotor(wiring []rune, notch rune) (*Rotor, error) {
	if len(wiring) != ALPHABET_SIZE {
		return nil, fmt.Errorf("invalid wiring length: %d", len(wiring))
	}

	if notch < 'A' || notch > 'Z' {
		return nil, fmt.Errorf("invalid notch: %c", notch)
	}

	r := &Rotor{
		wiring:   wiring,
		notch:    runeToAlphabetIndex(notch),
		position: 0,
	}

	return r, nil
}

func (r *Rotor) Notch() int {
	return r.notch
}

func (r *Rotor) Position() int {
	return r.position
}

func (r *Rotor) SetPosition(position int) error {
	if position < 0 || position > (ALPHABET_SIZE-1) {
		return fmt.Errorf("invalid position: %d", position)
	}
	r.position = position
	return nil
}

func (r *Rotor) TransformForward(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	index := (runeToAlphabetIndex(letter) + r.position) % ALPHABET_SIZE
	letter = r.wiring[index]
	return letter, nil
}

func (r *Rotor) TransformBackward(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	index := slices.Index(r.wiring, letter)
	if index == -1 {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	finalIndex := (index - r.position + ALPHABET_SIZE) % ALPHABET_SIZE
	return alphabetIndexToRune(finalIndex), nil
}

func (r *Rotor) Rotate() bool {
	rotateNext := r.position == r.notch
	r.position = (r.position + 1) % ALPHABET_SIZE
	return rotateNext
}
