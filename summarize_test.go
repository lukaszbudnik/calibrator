package calibrator

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummarize(t *testing.T) {
	file, err := os.Open("./test/stability.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	grouped, err := Group(records, "group")
	assert.Nil(t, err)

	dimension := "value"
	summary, err := Summarize(grouped["g1"], dimension)
	assert.Nil(t, err)
	assert.Equal(t, 6, summary.Samples)
}
