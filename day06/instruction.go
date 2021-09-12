package main

import (
	"strconv"
	"strings"
)

type operation int

const (
	TURN_ON operation = iota
	TURN_OFF
	TOGGLE
)

type instruction struct {
	op     operation
	startX int
	startY int
	endX   int
	endY   int
}

func parseInstruction(s string) instruction {
	var op operation
	var startRange, endRange []string
	var startX, startY, endX, endY int

	parts := strings.Split(s, " ")
	switch parts[0] {
	case "toggle":
		op = TOGGLE
		startRange = strings.Split(parts[1], ",")
		endRange = strings.Split(parts[3], ",")
	default:
		if parts[1] == "on" {
			op = TURN_ON
		} else {
			op = TURN_OFF
		}
		startRange = strings.Split(parts[2], ",")
		endRange = strings.Split(parts[4], ",")
	}

	startX, _ = strconv.Atoi(startRange[0])
	startY, _ = strconv.Atoi(startRange[1])
	endX, _ = strconv.Atoi(endRange[0])
	endY, _ = strconv.Atoi(endRange[1])

	return instruction{op, startX, startY, endX, endY}
}
