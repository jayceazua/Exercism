package scrabble

import "strings"

/*
Credit Cite: https://devhints.io/go, Edwin Cloud helped me unblock myself.
Letter                                                 Value
"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"       1
"D", "G"                                               2
"B", "C", "M", "P"                                     3
"F", "H", "V", "W", "Y"                                4
"K"                                                    5
"J", "X"                                               8
"Q", "Z"                                               10
*/

/*
// Score is a function that calculates the entire score point of a word.
func Score(word string) int {
	wu := strings.ToUpper(word)
	totalPoints := 0
	// setting variables to know their score <-- is not golang convention: use map
	scoreBoard := make(map[int][]string)
	// [(value, [letters here]), ...] <--- map the data
	scoreBoard[1] = []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	scoreBoard[2] = []string{"D", "G"}
	scoreBoard[3] = []string{"B", "C", "M", "P"}
	scoreBoard[4] = []string{"F", "H", "V", "W", "Y"}
	scoreBoard[5] = []string{"K"}
	scoreBoard[8] = []string{"J", "X"}
	scoreBoard[10] = []string{"Q", "Z"}

	for point, letters := range scoreBoard {
		// for loop of the word and getting individual letters
		for _, letter := range letters {
			for i := range wu {
				// if statements to match their points
				if string(wu[i]) == letter {
					// increment points
					totalPoints += point
				}
			}
		}
	}
	// return total
	return totalPoints
}
*/

// Score is a function that calculates the entire score point of a word.
func Score(word string) int {
	// create my map of the scoreBoard
	scoreBoard := make(map[string]int)
	scoreBoard["A"] = 1
	scoreBoard["E"] = 1
	scoreBoard["I"] = 1
	scoreBoard["O"] = 1
	scoreBoard["U"] = 1
	scoreBoard["L"] = 1
	scoreBoard["N"] = 1
	scoreBoard["R"] = 1
	scoreBoard["S"] = 1
	scoreBoard["T"] = 1
	scoreBoard["D"] = 2
	scoreBoard["G"] = 2
	scoreBoard["B"] = 3
	scoreBoard["C"] = 3
	scoreBoard["M"] = 3
	scoreBoard["P"] = 3
	scoreBoard["F"] = 4
	scoreBoard["H"] = 4
	scoreBoard["V"] = 4
	scoreBoard["W"] = 4
	scoreBoard["Y"] = 4
	scoreBoard["K"] = 5
	scoreBoard["J"] = 8
	scoreBoard["X"] = 8
	scoreBoard["Q"] = 10
	scoreBoard["Z"] = 10
	wu := strings.ToUpper(word)
	totalPoints := 0

	for i := range wu {
		l := string(wu[i])
		totalPoints += scoreBoard[l]
	}

	return totalPoints

}
