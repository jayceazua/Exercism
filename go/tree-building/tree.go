package tree

import (
	"errors"
	"sort"
)

// Record abstract datatype
type Record struct {
	ID, Parent int
}

// Node used to build a tree to model
type Node struct {
	ID       int
	Children []*Node
}

// Build create a tree of Nodes with given records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Parent < records[j].Parent || records[i].ID < records[j].ID
	})
	if (records[0] != Record{}) {
		return nil, errors.New("nada mudafakas")
	}

	lengthID := len(records)
	var result = Node{}
	nodeStorage := map[int]*Node{
		0: &result,
	}
	for _, record := range records[1:] {
		if _, exists := nodeStorage[record.ID]; exists {
			return nil, errors.New("madafaka is here already")
		}
		if record.ID <= record.Parent {
			return nil, errors.New("lil madafaka is less than big madafaka maybe equal")
		}
		if record.ID >= lengthID {
			return nil, errors.New("Non-continuous ID")
		}

		if parent, exists := nodeStorage[record.Parent]; exists {
			nn := Node{ID: record.ID}
			parent.Children = append(parent.Children, &nn)
			nodeStorage[record.ID] = &nn
		} else {
			return nil, errors.New("big madafaka not here")
		}
	}

	return &result, nil
}
