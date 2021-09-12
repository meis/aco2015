package main

import (
	"fmt"
	"strconv"
	"strings"
)

type board struct {
	connections map[string]string
	signals     map[string]uint16
}

func NewBoard(instructions []string) board {
	b := board{}
	b.connections = make(map[string]string, len(instructions))
	b.signals = make(map[string]uint16, len(instructions))

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " -> ")
		b.SetConnection(parts[1], parts[0])
	}

	return b
}

func (b *board) Run() {
	for k := range b.connections {
		b.GetSignal(k)
	}
}

func (b *board) GetSignal(id string) uint16 {
	parsedValue, err := strconv.ParseUint(id, 10, 16)
	if err == nil {
		return uint16(parsedValue)
	}
	_, keyFound := b.connections[id]
	if !keyFound {
		panic(fmt.Sprintf("Wire not found: %s", id))
	}
	signal, signalFound := b.signals[id]
	if signalFound {
		return signal
	}

	parts := strings.Split(b.connections[id], " ")
	switch {
	case len(parts) == 1:
		b.signals[id] = b.GetSignal(parts[0])
	case parts[0] == "NOT":
		b.signals[id] = ^b.GetSignal(parts[1])
	case parts[1] == "OR":
		b.signals[id] = b.GetSignal(parts[0]) | b.GetSignal(parts[2])
	case parts[1] == "AND":
		b.signals[id] = b.GetSignal(parts[0]) & b.GetSignal(parts[2])
	case parts[1] == "RSHIFT":
		b.signals[id] = b.GetSignal(parts[0]) >> b.GetSignal(parts[2])
	case parts[1] == "LSHIFT":
		b.signals[id] = b.GetSignal(parts[0]) << b.GetSignal(parts[2])
	default:
		panic(parts)
	}

	return b.signals[id]
}

func (b *board) SetConnection(k string, v string) {
	b.connections[k] = v
}
