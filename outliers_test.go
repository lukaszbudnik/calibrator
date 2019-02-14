package calibrator

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutliers(t *testing.T) {
	file, err := os.Open("./test/stability.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	// outliers for all records
	outliers, err := Outliers(records, "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 11)

	// outliers grouped
	grouped, err := Group(records, "group")
	assert.Nil(t, err)

	outliers, err = Outliers(grouped["g1"], "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 1)

	outliers, err = Outliers(grouped["g2"], "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 5)

	outliers, err = Outliers(grouped["g3"], "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 2)

	outliers, err = Outliers(grouped["g4"], "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 3)

	outliers, err = Outliers(grouped["g5"], "value")
	assert.Nil(t, err)
	assert.Len(t, outliers, 1)
}
