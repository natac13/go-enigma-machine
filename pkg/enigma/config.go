package enigma

const (
	BASE_ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ALPHABET_SIZE = 26

	ROTOR_I_WIRING = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	ROTOR_I_NOTCH  = 'Q'

	ROTOR_II_WIRING = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	ROTOR_II_NOTCH  = 'E'

	ROTOR_III_WIRING = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	ROTOR_III_NOTCH  = 'V'

	ROTOR_IV_WIRING = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
	ROTOR_IV_NOTCH  = 'J'

	ROTOR_V_WIRING = "VZBRGITYUPSDNHLXAWMJQOFECK"
	ROTOR_V_NOTCH  = 'Z'

	REFLECTOR_A_WIRING = "EJMZALYXVBWFCRQUONTSPIKHGD"
	REFLECTOR_B_WIRING = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	REFLECTOR_C_WIRING = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
)

func CreateRotorI() (*Rotor, error) {
	wiring := []rune(ROTOR_I_WIRING)
	notch := ROTOR_I_NOTCH
	return NewRotor(wiring, notch)
}

func CreateRotorII() (*Rotor, error) {
	wiring := []rune(ROTOR_II_WIRING)
	notch := ROTOR_II_NOTCH
	return NewRotor(wiring, notch)
}

func CreateRotorIII() (*Rotor, error) {
	wiring := []rune(ROTOR_III_WIRING)
	notch := ROTOR_III_NOTCH
	return NewRotor(wiring, notch)
}

func CreateRotorIV() (*Rotor, error) {
	wiring := []rune(ROTOR_IV_WIRING)
	notch := ROTOR_IV_NOTCH
	return NewRotor(wiring, notch)
}

func CreateRotorV() (*Rotor, error) {
	wiring := []rune(ROTOR_V_WIRING)
	notch := ROTOR_V_NOTCH
	return NewRotor(wiring, notch)
}

func CreateReflectorA() (*Reflector, error) {
	wiring := []rune(REFLECTOR_A_WIRING)
	return NewReflector(wiring)
}

func CreateReflectorB() (*Reflector, error) {
	wiring := []rune(REFLECTOR_B_WIRING)
	return NewReflector(wiring)
}

func CreateReflectorC() (*Reflector, error) {
	wiring := []rune(REFLECTOR_C_WIRING)
	return NewReflector(wiring)
}
