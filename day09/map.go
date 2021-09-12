package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Map struct {
	distances map[string]map[string]int
	cities    map[string]bool
}

func NewMap(distances []string) Map {
	m := Map{make(map[string]map[string]int), make(map[string]bool)}

	for _, s := range distances {
		parts := strings.Split(s, " ")
		d, _ := strconv.Atoi(parts[4])
		m.AddDistance(parts[0], parts[2], d)
	}

	return m
}

func (m *Map) AddDistance(origin string, destiny string, distance int) {
	_, exists := m.distances[origin]
	if !exists {
		m.distances[origin] = make(map[string]int)
	}
	m.distances[origin][destiny] = distance
	m.cities[origin] = true
	m.cities[destiny] = true
}

func (m *Map) GetDistance(origin string, destiny string) int {
	if origin == destiny {
		panic(fmt.Sprintf("Asking for distance between %s -> %s", origin, destiny))
	}

	distance, exists := m.distances[origin][destiny]
	if exists {
		return distance
	}

	distance, exists = m.distances[destiny][origin]
	if exists {
		return distance
	}

	panic(fmt.Sprintf("No entry for %s -> %s", origin, destiny))
}

func (m *Map) GetCities() []string {
	cities := make([]string, 0, len(m.cities))

	for k, _ := range m.cities {
		cities = append(cities, k)
	}

	return cities
}
