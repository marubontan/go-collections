package collections

import (
	"errors"
	"sort"
)

type Counter[T comparable] struct {
	Data map[T]int
}

func NewCounter[T comparable](data []T) (Counter[T], error) {
	if len(data) == 0 {
		return Counter[T]{Data: make(map[T]int)}, errors.New("empty data")

	}
	count := make(map[T]int)
	for _, v := range data {
		count[v] += 1
	}
	return Counter[T]{Data: count}, nil
}

func (c *Counter[T]) MostCommon(n int) []map[T]int {
	keys := make([]T, 0, len(c.Data))
	for k, _ := range c.Data {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return c.Data[keys[i]] > c.Data[keys[j]]
	})
	commonItems := make([]map[T]int, 0, n)
	for _, key := range keys[:n] {
		commonItems = append(commonItems, map[T]int{key: c.Data[key]})
	}
	return commonItems
}
