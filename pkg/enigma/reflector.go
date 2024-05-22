package enigma

import "fmt"

type Reflector struct {
	wiring []rune
}

func NewReflector(wiring []rune) (*Reflector, error) {
	if len(wiring) != ALPHABET_SIZE {
		return nil, fmt.Errorf("invalid wiring length: %d", len(wiring))
	}

	r := &Reflector{
		wiring: wiring,
	}

	return r, nil
}

func (r *Reflector) Transform(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}

	return r.wiring[runeToAlphabetIndex(letter)], nil
}
