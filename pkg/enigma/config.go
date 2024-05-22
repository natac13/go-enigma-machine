package enigma

const (
	BASE_ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ALPHABET_SIZE = 26

	ROTOR_I_WIRING   = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	ROTOR_I_NOTCH    = 'Q'
	ROTOR_II_WIRING  = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	ROTOR_II_NOTCH   = 'E'
	ROTOR_III_WIRING = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	ROTOR_III_NOTCH  = 'V'

	REFLECTOR_B_WIRING = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	REFLECTOR_C_WIRING = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
)

type EnigmaConfig struct {
	Rotors []*Rotor
	// reflector *Reflector
	Plugboard            *Plugboard
	RotorPositions       []int
	RotorOrder           []int
	RotorRingSettings    []int
	PlugboardConnections map[rune]rune
}

func NewEnigmaConfig() *EnigmaConfig {
	return &EnigmaConfig{
		Rotors:               []*Rotor{},
		Plugboard:            NewPlugboard(),
		RotorPositions:       []int{},
		RotorOrder:           []int{},
		RotorRingSettings:    []int{},
		PlugboardConnections: map[rune]rune{},
	}
}
