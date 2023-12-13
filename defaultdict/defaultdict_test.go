package defaultdict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntIntDefaultDict(t *testing.T) {
	defaultDict := NewDefaultDict[int, int]()
	defaultDict.Set(1, 2)
	assert.Equal(t, 2, defaultDict.Get(1))
	assert.Equal(t, 0, defaultDict.Get(2))
}

func TestStrStrDefaultDict(t *testing.T) {
	defaultDict := NewDefaultDict[string, string]()
	defaultDict.Set("1", "2")
	assert.Equal(t, "2", defaultDict.Get("1"))
	assert.Equal(t, "", defaultDict.Get("0"))
}

func TestStrSliceDefaultDict(t *testing.T) {
	defaultDict := NewDefaultDict[string, []int]()
	defaultDict.Set("1", []int{1})
	assert.Equal(t, []int{1}, defaultDict.Get("1"))
	assert.Equal(t, []int(nil), defaultDict.Get("0"))
	defaultDict.data["1"] = append(defaultDict.Get("1"), 2)
	assert.Equal(t, []int{1, 2}, defaultDict.Get("1"))
}

func TestStrMapDefaultDict(t *testing.T) {
	defaultDict := NewDefaultDict[string, map[string]int]()
	defaultDict.Set("1", map[string]int{"1": 1})
	assert.Equal(t, map[string]int{"1": 1}, defaultDict.Get("1"))
	assert.Equal(t, map[string]int(nil), defaultDict.Get("0"))
}
