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

	/*
	 *  If you wanna delete certain element from a map we can use delete()
	 *  delete( <Name of the map>, <key>)
	 */

	delete(dayMonths, "Feb")

	// If we try to delete a element which is not in the map, then go wont complain. no error.
	delete(dayMonths, "February")

	has28 := 0

	for _, days := range dayMonths {
		if days == 28 {
			has28 += 1
		}

	}

	fmt.Printf("%d months have 28 days \n", has28)

}
