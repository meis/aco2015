package main

import (
	"fmt"
	"math"

	"github.com/meis/aoc2015/input"
)

func main() {
	distances := input.GetInputStrings()

	fmt.Printf("The shortest distance is: %d\n", part1(distances))
	fmt.Printf("The longest distance is: %d\n", part2(distances))
}

func part1(distances []string) int {
	m := NewMap(distances)

	return recursiveMinDistance(m, "", m.GetCities())
}

func part2(distances []string) int {
	m := NewMap(distances)

	return recursiveMaxDistance(m, "", m.GetCities())
}

func recursiveMinDistance(m Map, city string, citiesLeft []string) int {
	totalDistance := 0

	if len(citiesLeft) == 1 {
		return m.GetDistance(city, citiesLeft[0])
	} else {

		minDistance := math.MaxInt64
		for i, nextCity := range citiesLeft {
			distance := recursiveMinDistance(m, nextCity, withIndexRemoved(citiesLeft, i))
			if city != "" {
				distance += m.GetDistance(city, nextCity)
			}

			if distance < minDistance {
				minDistance = distance
			}
		}

		totalDistance += minDistance
	}

	return totalDistance
}

func recursiveMaxDistance(m Map, city string, citiesLeft []string) int {
	totalDistance := 0

	if len(citiesLeft) == 1 {
		return m.GetDistance(city, citiesLeft[0])
	} else {

		maxDistance := 0
		for i, nextCity := range citiesLeft {
			distance := recursiveMaxDistance(m, nextCity, withIndexRemoved(citiesLeft, i))
			if city != "" {
				distance += m.GetDistance(city, nextCity)
			}

			if distance > maxDistance {
				maxDistance = distance
			}
		}

		totalDistance += maxDistance
	}

	return totalDistance
}

func withIndexRemoved(slice []string, index int) []string {
	newSlice := make([]string, len(slice))
	copy(newSlice, slice)

	// Delete element: replace the one to replace with the last one and remove the last one
	newSlice[index] = newSlice[len(newSlice)-1]
	return newSlice[:len(newSlice)-1]
}
