package enigma

import (
	"fmt"
)

type Plugboard struct {
	connections map[rune]rune
}

func NewPlugboard() *Plugboard {
	return &Plugboard{connections: map[rune]rune{}}
}

func (p *Plugboard) addConnection(a, b rune) error {
	if a < 'A' || a > 'Z' || b < 'A' || b > 'Z' {
		return fmt.Errorf("invalid connection: %c %c", a, b)
	}

	if a == b {
		return fmt.Errorf("cannot connect a letter to itself: %c %c", a, b)
	}

	if _, ok := p.connections[a]; ok {
		return fmt.Errorf("letter %c is already connected", a)
	}

	if _, ok := p.connections[b]; ok {
		return fmt.Errorf("letter %c is already connected", b)
	}

	if p.countConnections() == 10 {
		return fmt.Errorf("cannot add more than 10 connections")
	}

	p.connections[a] = b
	p.connections[b] = a

	return nil
}

func (p *Plugboard) removeConnection(a rune) error {
	if a < 'A' || a > 'Z' {
		return fmt.Errorf("invalid connection: %c", a)
	}

	if _, ok := p.connections[a]; !ok {
		return fmt.Errorf("letter %c is not connected", a)
	}

	delete(p.connections, p.connections[a])
	delete(p.connections, a)

	return nil
}

func (p *Plugboard) clearConnections() {
	p.connections = map[rune]rune{}
}

func (p *Plugboard) transform(letter rune) (rune, error) {
	if letter < 'A' || letter > 'Z' {
		return 0, fmt.Errorf("invalid letter: %c", letter)
	}
	if connection, ok := p.connections[letter]; ok {
		return connection, nil
	}
	return letter, nil
}

func (p *Plugboard) countConnections() int {
	count := len(p.connections)
	if count == 0 {
		return 0
	}
	return count / 2
}

// func (p *Plugboard) String() string {
// 	var connections []string
// 	for a, b := range p.connections {
// 		if a < b {
// 			connections = append(connections, fmt.Sprintf("%c%c", a, b))
// 		}
// 	}
// 	return fmt.Sprintf("Plugboard{%s}", connections)
// }
