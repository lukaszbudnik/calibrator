package calibrator

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaders(t *testing.T) {
	file, err := os.Open("./test/stability.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	// leaders for all records
	leaders, err := Leaders(records, "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 21)

	// leaders grouped
	grouped, err := Group(records, "group")
	assert.Nil(t, err)

	leaders, err = Leaders(grouped["g1"], "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 3)

	leaders, err = Leaders(grouped["g2"], "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 12)

	leaders, err = Leaders(grouped["g3"], "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 5)

	leaders, err = Leaders(grouped["g4"], "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 4)

	leaders, err = Leaders(grouped["g5"], "value")
	assert.Nil(t, err)
	assert.Len(t, leaders, 2)
}
