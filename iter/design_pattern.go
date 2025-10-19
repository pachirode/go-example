package main

type Iterator struct {
	data  []int
	index int
}

func NewIterator(data []int) *Iterator {
	return &Iterator{data: data, index: 0}
}

func (it *Iterator) HasNext() bool {
	return it.index < len(it.data)
}

func (it *Iterator) Next() int {
	if it.HasNext() {
		panic("Stop iteration")
	}
	value := it.data[it.index]
	it.index++
	return value
}
