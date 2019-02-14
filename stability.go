package calibrator

import (
	"errors"
	"fmt"

	"github.com/montanaflynn/stats"
)

func Stability(records [][]string, dimension string, threshold float64) (bool, error) {

	if len(dimension) == 0 {
		return false, errors.New("Dimension must not be empty")
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
		return false, fmt.Errorf("Dimension not found: %v", dimension)
	}

	values := []string{}
	for _, record := range records[1:len(records)] {
		values = append(values, record[column])
	}
	data := stats.LoadRawData(values)

	variance, err := data.Variance()
	if err != nil {
		return false, err
	}

	if variance < threshold {
		return true, nil
	}
	return false, nil
}
