package robotname

import (
	"math/rand"
	"time"
)

// Robot struct type to model our robot's name
type Robot struct {
	name string
}

var nameBank = map[string]bool{}

// Name the robot is no name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		// call the getName function 
		r.name = getName()
	}
	return r.name, nil
}

// Reset takes the rebot and resets all its BIOS
func (r *Robot) Reset() {
	r.name = ""
}

// getName generates a random name for the robot as well checks if the name is in the nameBank.
func getName() string {
	name := ""
	bs := [5]byte{}
	// sets us up for true random true ints; little trick learned from CS1.2
	rand.Seed(time.Now().UnixNano())

	for {
		rand.Read(bs[:])
		// generates random letters
		bs[0] = bs[0]%26 + 'A'
		bs[1] = bs[1]%26 + 'A'
		// generates random integers: below
		bs[2] = bs[2]%10 + '0'
		bs[3] = bs[3]%10 + '0'
		bs[4] = bs[4]%10 + '0'

		name = string(bs[:])
		if _, n := nameBank[name]; !n {
			nameBank[name] = true
			break
		}
	}
	return name
}
