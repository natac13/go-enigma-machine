package enigma

import (
	"fmt"
	"slices"
	"strings"
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

// setPosition sets the rotor position based on a letter.
// the letter must be a single letter from A to Z.
// the position is the zero-based index of the letter in the alphabet.
func (r *Rotor) setPosition(letter string) error {
	if len(letter) != 1 {
		return fmt.Errorf("invalid letter: %s", letter)
	}

	normalized := strings.ToUpper(letter)
	lr := rune(normalized[0])
	if lr < 'A' || lr > 'Z' {
		return fmt.Errorf("invalid letter: %c", lr)
	}

	r.position = runeToAlphabetIndex(lr)
	return nil
}

// rotate returns true if the rotor should rotate the next rotor
func (r *Rotor) rotate() bool {
	rotateNext := r.position == r.notch
	r.position = (r.position + 1) % ALPHABET_SIZE
	return rotateNext
}

// transformForward transforms a letter through the rotor from right to left
func (r *Rotor) transformForward(letter rune) (rune, error) {
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

// transformBackward transforms a letter through the rotor from left to right
func (r *Rotor) transformBackward(letter rune) (rune, error) {
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
