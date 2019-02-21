package twelve

import "fmt"

var days = [12]string{"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}
var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Verse produces the correct verse needed for the index
func Verse(day int) string {
	d := days[day-1]
	var g string
	if day == 1 {
		g = gifts[0]
	} else {
		for i := day - 1; i > 0; i-- {
			g += gifts[i] + ", "
		}
		g += "and " + gifts[0]
	}

	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", d, g)
	return verse
}

// Song returns the lyrics in chronological order
func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}