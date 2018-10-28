package map_ranker

import "github.com/dinocorpllc/golang-perf"

type MapRanker struct {
	data map[string]uint64
}

func (m *MapRanker) TopN(n uint64) []uint64 {
	acc := ranker.NewSliceAccumulator(n)
	for _, v := range m.data {
		acc.Offer(v)
	}
	return acc.Get()
}

func (m *MapRanker) Insert(key string, value uint64) {
	m.data[key] = value
}

func (m *MapRanker) Get(key string) uint64 {
	v, _ := m.data[key]
	return v
}

func (m *MapRanker) Delete(key string) bool {
	delete(m.data, key)
	return true
}