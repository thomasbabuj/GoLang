package main

import (
	"fmt"
)

// Maps examples
/*
   - Maps have index and it has corresponding values associated with this
   - Index can be any in type of value
   - We can use make() to create a new map
*/

func main() {

	// Allocating the values which we already know for the map can be done like,

	dayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	has28 := 0

	for _, days := range dayMonths {
		if days == 28 {
			has28 += 1
		}

	}

	fmt.Printf("%d months have 28 days \n", has28)

}
