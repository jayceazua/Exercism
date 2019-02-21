package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

/*
Notes: https://golang.org/pkg/bufio/#NewScanner
Package bufio implements buffered I/O. 
It wraps an io.Reader or io.Writer object, 
creating another object (Reader or Writer) that also 
implements the interface but provides buffering and some help for textual I/O.
*/

// Team struct decribes the team blueprint
type Team struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

// Tally tallies football matches.
func Tally(r io.Reader, w io.Writer) error {
	// NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines.
	s := bufio.NewScanner(r)

	tms := map[string]*Team{}
	var teams []*Team

	for s.Scan() {
		l := s.Text()

		if len(l) == 0 || strings.HasPrefix(l, "#") {
			continue
		}

		fields := strings.Split(l, ";")

		if len(fields) != 3 {
			return fmt.Errorf("got todo better than that")
		}

		if _, ok := tms[fields[0]]; !ok {
			t := Team{name: fields[0]}
			tms[fields[0]] = &t
			teams = append(teams, &t)
		}

		if _, ok := tms[fields[1]]; !ok {
			t := Team{name: fields[1]}
			tms[fields[1]] = &t
			teams = append(teams, &t)
		}

		switch fields[2] {
			case "win":
				tms[fields[0]].win()
				tms[fields[1]].loss()
			case "draw":
				tms[fields[0]].draw()
				tms[fields[1]].draw()
			case "loss":
				tms[fields[0]].loss()
				tms[fields[1]].win()
			default:
				return fmt.Errorf("got todo better than that")
		}

		sort.Slice(teams, func(i, j int) bool {
			if teams[i].points == teams[j].points {
				return teams[i].name < teams[j].name
			}
			return teams[i].points > teams[j].points
		})


	}

	if _, err := fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P"); err != nil {
		return err
	}

	for _, t := range teams {
		if _, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", t.name, t.matches, t.wins, t.draws, t.losses, t.points); err != nil {
			return err
		}
	}

	return nil
}

// Outcomes from the matches
// win gives the winning team 3 points
func (t *Team) win() {
	t.matches++
	t.wins++
	t.points += 3
}

// draw gives each team 1 point
func (t *Team) draw() {
	t.matches++
	t.draws++
	t.points++
}

// loss gives the team no point
func (t *Team) loss() {
	t.matches++
	t.losses++
}
