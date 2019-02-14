package luhn

import (
	"strconv"
	"strings"
)

// Valid takes in an input of a string number and checks if it passes the luhn formula
func Valid(n string) bool {
	var ts, p int
	// clean up input
	n = strings.Replace(n, " ", "", -1)
	// cancel out if less than 1
	if len(n) <= 1 {
		return false
	}

	// adding all the integers with the double
	for i := len(n) - 1; i >= 0; i-- {
		// convert string to rune to integer
		d, e := strconv.Atoi(string(n[i]))
		// if the convert fails return false
		if e != nil {
			return false
		}

		p++ // keeping count of the position I am in
		// if the integer is the second position from right
		if p%2 == 0 {
			// multiply it by 2
			d *= 2
			// if the digit is greater than 9
			if d > 9 {
				// substract 9 from it
				d -= 9
			}

		}
		// add the digit to the total sum
		ts += d
	}
	// mod 10 check; has to equal to 0
	return ts%10 == 0

}
