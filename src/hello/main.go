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
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	// Simple Map usage
	fmt.Printf("Days in February : %d \n", dayMonths["Feb"])
}
