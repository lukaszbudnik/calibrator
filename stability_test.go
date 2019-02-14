package calibrator

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStability(t *testing.T) {
	file, err := os.Open("./test/stability.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	grouped, err := Group(records, "group")
	assert.Nil(t, err)

	// g1 and g3 are stable
	stable, err := Stability(grouped["g1"], "value", 10)
	assert.Nil(t, err)
	assert.True(t, stable)

	stable, err = Stability(grouped["g3"], "value", 10)
	assert.Nil(t, err)
	assert.True(t, stable)

	// g2, g4, and g5 unstable
	stable, err = Stability(grouped["g2"], "value", 10)
	assert.Nil(t, err)
	assert.False(t, stable)
	stable, err = Stability(grouped["g4"], "value", 10)
	assert.Nil(t, err)
	assert.False(t, stable)
	stable, err = Stability(grouped["g5"], "value", 10)
	assert.Nil(t, err)
	assert.False(t, stable)
}
