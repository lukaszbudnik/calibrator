package calibrator

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	file, err := os.Open("./test/data.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	assert.Len(t, records, 6)

	dimension := "group"
	grouped, err := Group(records, dimension)
	assert.Nil(t, err)
	assert.Len(t, grouped, 2)

	dimension = "subgroup"
	grouped, err = Group(records, dimension)
	assert.Nil(t, err)
	assert.Len(t, grouped, 3)
}

func TestRecursiveGroup(t *testing.T) {
	file, err := os.Open("./test/data.csv")
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	data := csv.NewReader(reader)
	records, err := data.ReadAll()
	assert.Nil(t, err)

	assert.Len(t, records, 6)

	results, _ := RecursiveGroup(records, []string{"group", "subgroup"})

	// top level groups should be home and work
	assert.Len(t, results, 2)

	var work *RecGroup
	var home *RecGroup
	for _, group := range results {
		if group.Dimension == "work" {
			work = group
		}
		if group.Dimension == "home" {
			home = group
		}
	}

	assert.NotNil(t, work)
	assert.NotNil(t, home)

	// under home group there are 3 subgroups
	assert.Len(t, home.Subgroups, 3)
	assert.Len(t, home.Records, 5)
	// again driven by maps order may vary, sort results
	expected := []string{"temp", "humidity", "light"}
	actual := []string{home.Subgroups[0].Dimension, home.Subgroups[1].Dimension, home.Subgroups[2].Dimension}
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	// under work group there is 1 subgroup
	assert.Len(t, work.Subgroups, 1)
	assert.Len(t, work.Records, 2)
	assert.Equal(t, "temp", work.Subgroups[0].Dimension)

	// sum all records should be 9 (each tree leaf contains headers)
	assert.Equal(t, 9, len(home.Subgroups[0].Records)+len(home.Subgroups[1].Records)+len(home.Subgroups[2].Records)+len(work.Subgroups[0].Records))

	// printGroup(results)
}

func printGroup(grps []*RecGroup) {
	for _, grp := range grps {
		log.Printf("%v", grp.Dimension)
		printGroup(grp.Subgroups)
		if len(grp.Records) > 0 {
			log.Println(grp.Records)
		}
	}
}
