package calibrator

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	file, err := os.Open("./test/data.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	condition := map[string]string{"group": "home"}
	filtered, err := Filter(records, condition)
	assert.Nil(t, err)
	// filter preserves header
	assert.Len(t, filtered, 5)

	condition = map[string]string{"subgroup": "temp"}
	filtered, err = Filter(records, condition)
	assert.Nil(t, err)
	// filter preserves header
	assert.Len(t, filtered, 4)

	condition = map[string]string{"group": "home", "subgroup": "temp"}
	filtered, err = Filter(records, condition)
	assert.Nil(t, err)
	// filter preserves header
	assert.Len(t, filtered, 3)
}
