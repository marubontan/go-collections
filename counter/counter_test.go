package collections

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntCounter(t *testing.T) {
	counter, _ := NewCounter([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5})
	assert.Equal(t, []map[int]int{{5: 5}, {4: 4}}, counter.MostCommon(2))
}

func TestStringCounter(t *testing.T) {
	counter, _ := NewCounter([]string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4", "5", "5", "5", "5", "5"})
	assert.Equal(t, []map[string]int{{"5": 5}, {"4": 4}}, counter.MostCommon(2))
}

func TestEmptyCounter(t *testing.T) {
	_, err := NewCounter([]int{})
	assert.Equal(t, err, errors.New("empty data"))

}
