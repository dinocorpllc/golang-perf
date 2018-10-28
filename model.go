package ranker

// CRUD defines the interface of our synthetic workload.
type Ranker interface {

	// Iterate executes do(...) on each value in the collection
	TopN(n uint64) []uint64

	// Insert adds a k,v pair to the collection
	Insert(key string, value uint64)

	// Read gets the value at key
	Get(key string) uint64

	// Delete deletes the value at key
	Delete(key string) bool
}

// Accumulator defines the interface for data structure that accumulates the top N scores offered to it
type TopNAccumulator interface {
	Offer(n uint64)
	Get() []uint64
}

type sliceAccumulator struct {
	vals []uint64
	n uint64
}

func NewSliceAccumulator(n uint64) *sliceAccumulator {
	return &sliceAccumulator{make([]uint64, n, n), n}
}

func (s *sliceAccumulator) Offer(val uint64) {
	if val <= s.vals[s.n-1] {
		return
	}
	var i int
	var cur uint64
	for i, cur = range s.vals {
		if val > cur {
			break
		}
	}
	tail := append([]uint64{val}, s.vals[i:(s.n - 1)]...)
	s.vals = append(s.vals[0:i], tail...)
}

func (s *sliceAccumulator) Get() []uint64 {
	return s.vals
}


