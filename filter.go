package calibrator

import (
	"errors"
	"fmt"
)

func Filter(records [][]string, condition map[string]string) ([][]string, error) {
	if condition == nil || len(condition) == 0 {
		return nil, errors.New("Condition must not be nil or empty")
	}
	columns := map[int]string{}
	header := records[0]
	for dimension := range condition {
		for i, column := range header {
			if column == dimension {
				columns[i] = condition[dimension]
				break
			}
		}
	}

	if len(columns) < len(condition) {
		return nil, fmt.Errorf("Dimension(s) not found in condition: %v", (len(condition) - len(columns)))
	}

	filtered := [][]string{}
	filtered = append(filtered, header)
	for _, record := range records[1:len(records)] {
		match := false
		for i := range columns {
			if record[i] == columns[i] {
				match = true
			} else {
				match = false
				break
			}
		}
		if match {
			filtered = append(filtered, record)
		}
	}

	return filtered, nil
}
