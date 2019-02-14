package grains

import (
	"errors"
)

// Square returns the number of grains on any given chessboard square
func Square(s int) (uint64, error) {
	// Solve the edge case
	if s < 1 || s > 64 {
		return 0, errors.New("invalid chessboard")
	}
	// math.Pow(2, float64(/*this is the given square*/i-1/*the minue one gives it reason to be accurate*/)) <- did not help because it did not return the integer uint64
	return 1 << uint64(s-1), nil
}

// Total returns the accumulated number of grains at a given chessboard square
func Total() uint64 {
	return 1<<64 - 1
}

/*
https://www.tutorialspoint.com/go/go_operators.htm
Used this link to help me convert floats into a uint64
The << and >> operator does this for you.
<< Binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.
>> Binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand.
*/
