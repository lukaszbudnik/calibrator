package calibrator

import (
	"errors"
	"fmt"

	"github.com/montanaflynn/stats"
)

type Summary struct {
	Samples   int             // contains number of sample
	Variance  float64         // variance of specified dimension
	Min       float64         // min value of the specified dimension
	Max       float64         // max value of the specified dimension
	Mean      float64         // mean value of the specified dimension
	Median    float64         // median value of the specified dimension
	Quartiles stats.Quartiles // Q1, Q2, and Q3 quartiles of the specified dimension
	Outliers  stats.Outliers  // outliers values of the specified dimension
}

func Summarize(records [][]string, dimension string) (Summary, error) {
	if len(dimension) == 0 {
		return Summary{}, errors.New("Dimension must not be empty")
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
		return Summary{}, fmt.Errorf("Dimension not found: %v", dimension)
	}

	values := []string{}
	for _, record := range records[1:len(records)] {
		values = append(values, record[column])
	}
	data := stats.LoadRawData(values)

	outliers, err := stats.QuartileOutliers(data)
	if err != nil {
		return Summary{}, err
	}

	quartiles, err := stats.Quartile(data)
	if err != nil {
		return Summary{}, err
	}

	min, err := data.Min()
	if err != nil {
		return Summary{}, err
	}

	max, err := data.Max()
	if err != nil {
		return Summary{}, err
	}

	mean, err := data.Mean()
	if err != nil {
		return Summary{}, err
	}

	median, err := data.Median()
	if err != nil {
		return Summary{}, err
	}

	variance, err := data.Variance()
	if err != nil {
		return Summary{}, err
	}

	return Summary{
		Samples:   len(records) - 1,
		Quartiles: quartiles,
		Outliers:  outliers,
		Min:       min,
		Max:       max,
		Mean:      mean,
		Median:    median,
		Variance:  variance,
	}, nil
}
