package calibrator

import (
	"errors"
	"fmt"
)

type RecGroup struct {
	Dimension string
	Records   [][]string
	Subgroups []*RecGroup
}

func Group(records [][]string, dimension string) (map[string][][]string, error) {
	if len(dimension) == 0 {
		return nil, errors.New("Dimension must not be empty")
	}

	column := -1
	header := records[0]
	for i, col := range header {
		if col == dimension {
			column = i
			break
		}
	}

	if column == -1 {
		return nil, fmt.Errorf("Dimension not found: %v", dimension)
	}

	grouped := make(map[string][][]string)
	for _, record := range records[1:len(records)] {
		key := record[column]
		group, ok := grouped[key]
		if !ok {
			// always copy header
			group = [][]string{header}
		}
		group = append(group, record)
		grouped[key] = group
	}

	return grouped, nil
}

func RecursiveGroup(records [][]string, dimensions []string) ([]*RecGroup, error) {
	if len(dimensions) == 0 {
		return nil, errors.New("Dimensions cannot be empty")
	}

	groups, err := Group(records, dimensions[0])
	if err != nil {
		return nil, err
	}

	current := []*RecGroup{}
	for key, group := range groups {
		grp := &RecGroup{Dimension: key}
		if len(dimensions) > 1 {
			subgroups, err := RecursiveGroup(group, dimensions[1:len(dimensions)])
			if err != nil {
				return nil, err
			}
			grp.Subgroups = subgroups
			grp.Records = group
		} else {
			grp.Records = group
		}
		current = append(current, grp)
	}
	return current, nil
}
