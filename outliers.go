package calibrator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/montanaflynn/stats"
)

func Outliers(records [][]string, dimension string) ([][]string, error) {
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

	values := []string{}
	for _, record := range records[1:len(records)] {
		values = append(values, record[column])
	}
	data := stats.LoadRawData(values)

	outliers, err := stats.QuartileOutliers(data)
	if err != nil {
		return nil, err
	}

	// copy header
	outlierRecords := [][]string{header}

r:
	for _, record := range records {
		for _, outlier := range outliers.Extreme {
			value, _ := strconv.ParseFloat(record[column], 10)
			if value == outlier {
				outlierRecords = append(outlierRecords, record)
				continue r
			}
		}
	}
	return outlierRecords, nil
}
