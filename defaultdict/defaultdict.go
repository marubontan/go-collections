package collections

type DefaultDict[K comparable, V any] struct {
	data map[K]V
}

func NewDefaultDict[K comparable, V any]() DefaultDict[K, V] {
	return DefaultDict[K, V]{data: make(map[K]V)}
}

func (d *DefaultDict[K, V]) Set(k K, v V) {
	d.data[k] = v
}

func (d *DefaultDict[K, V]) Get(k K) V {
	if v, ok := d.data[k]; ok {
		return v
	} else {
		var zeroValue V
		d.Set(k, zeroValue)
		return d.Get(k)
	}
}
