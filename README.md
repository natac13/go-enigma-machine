# Go Enigma Machine

This is a simple implementation of the Enigma Machine in Go. It is based on the Enigma I, the most common version of the machine used by the German military during World War II.

The Enigma Machine is a simple substitution cipher. It works by taking a letter of the alphabet and substituting it with another letter. The substitution is determined by the settings of the machine, which can be changed by the operator.

## Features

- Simulates the Enigma I Machine
- Configurable rotors, reflectors, rotor positions, rotor ring settings, and plugboard pairs
- Command-line interface
- Configurable settings using flags or a config file

## Prerequisites

- Go 1.22.2 or later

## Installation

To install the package, run the following command:

```bash
go install github.com/natac13/go-enigma-machine
```

You will then have an `go-enigma-machine` executable in your `$GOPATH/bin` directory.

## Usage

The CLI tool takes a message as input and outputs the encrypted message.
To use the tool, run the following command:

```bash
go-enigma-machine encrypt "bootdev rocks"
```

**Output**:

```plaintext
Enigma machine settings used:
- Reflector: B
- Rotors: [III II I]
- Rotor positions: AAA
- Rotor ring settings: AAA
- Plugboard pairs: []

Original message: bootdev rocks
Encrypted message: WLQUC DIFFV VH
```

### With Flags

You can also specify the settings of the Enigma Machine using flags.
For example:

```bash
go-enigma-machine encrypt "bootdev rocks" --reflector C --rotors III,IV,II --rotor-positions ABC --rotor-ring-settings DEF --plugboard-pairs AB,CD,EF
```

**Output**:

```plaintext
Enigma machine settings used:
- Reflector: C
- Rotors: [III IV II]
- Rotor positions: ABC
- Rotor ring settings: DEF
- Plugboard pairs: [AB CD EF]

Original message: bootdev rocks
Encrypted message: QKHYV RICZR BB
```

### With Config File

You can also specify the settings in a config file. The config file should be in YAML format.

The file location can be specified using the `--config` flag.
And by default, the CLI will look for a file named `.go-enigma-machine.yaml` in your home directory.

For example:

```yaml
rotors:
  - III # Leftmost rotor or first rotor
  - II
  - I # Rightmost rotor or last rotor
rotor-positions: AAA
rotor-ring-settings: AAA
reflector: B
plugboard:
  pairs:
    - AB
    - CD
    - EF
```

## Configuration Options

The following settings can be configured:

- **Reflector**: Choose from `A`, `B`, or `C`.
- **Rotors**: Choose from `I`, `II`, `III`, `IV`, or `V`.
- **Rotor Positions**: A three-letter string representing the initial position of the rotors. (e.g., `AAA`).
- **Rotor Ring Settings**: A three-letter string representing the initial ring setting of the rotors. (e.g., `AAA`).
- **Plugboard Pairs**: A list of pairs of letters that are swapped before and after the encryption process. (e.g., `AB,CD,EF`).

## Flags

The following flags can be used to configure the Enigma Machine:

- `--reflector` or `u`: Choose from `A`, `B`, or `C`.
- `--rotors` or `r`: A list of three rotors to use. (e.g., `I,II,III`). The leftmost rotor is the first rotor, and the rightmost rotor is the last rotor.
- `--rotor-positions` or `d`: A three-letter string representing the initial position of the rotors. (e.g., `AAA`).
- `--rotor-ring-settings` or `s`: A three-letter string representing the initial ring setting of the rotors. (e.g., `AAA`).
- `--plugboard-pairs` or `p`: A list of pairs of letters that are swapped before and after the encryption process. (e.g., `AB,CD,EF`).

**Fun Facts**:

- The `u` shorthand for reflector selection stand for [U]mkehrwalze, German for "reversing rotor".
- The `s` shorthand for rotor ring settings stand for Ring[s]tellung, German for "ring setting".

## Contributing

This was a project I did to learn more about Go.
If you have any suggestions or improvements, feel free to open an issue or a pull request.

## License

MIT
