
package raindrops
import "strconv"

func Convert(num int) string {
		result := ""
		if num % 3 == 0 {
			result += "Pling"
		}
		if num % 5 == 0 {
			result += "Plang"
		}
		if num % 7 == 0 {
			result += "Plong"
		}

		if result == "" {
			result += strconv.Itoa(num)
		}
		return result
}

/*
If the number has 3 as a factor, output 'Pling'.
If the number has 5 as a factor, output 'Plang'.
If the number has 7 as a factor, output 'Plong'.
Site Credit: https://stackoverflow.com/questions/10105935/how-to-convert-an-int-value-to-string-in-go
*/