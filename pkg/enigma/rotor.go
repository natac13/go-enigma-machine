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

func (r *Rotor) Rotate() bool {
	rotateNext := r.position == r.notch
	r.position = (r.position + 1) % ALPHABET_SIZE
	return rotateNext
}

func (r *Rotor) TransformForward(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	// find the index of the letter in the wiring based on the position
	index := (runeToAlphabetIndex(letter) + r.position) % ALPHABET_SIZE
	// find the connection in the wiring based on the index
	transformed := r.wiring[index]
	// find the index of the transformed letter in the alphabet
	finalIndex := (runeToAlphabetIndex(transformed) - r.position + ALPHABET_SIZE) % ALPHABET_SIZE
	// find the letter in the alphabet based on the final index
	transformed = alphabetIndexToRune(finalIndex)

	return transformed, nil
}

func (r *Rotor) TransformBackward(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	indexOfIncomingInAlphabet := runeToAlphabetIndex(letter)
	realIndex := (indexOfIncomingInAlphabet + r.position) % ALPHABET_SIZE
	realLetter := alphabetIndexToRune(realIndex)

	index := slices.Index(r.wiring, realLetter)
	if index == -1 {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	transformedIndex := (index - r.position + ALPHABET_SIZE) % ALPHABET_SIZE
	transformed := alphabetIndexToRune(transformedIndex)
	return transformed, nil
}
