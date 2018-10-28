package presorted_ranker

import (
	_ "fmt"
	"strings"
)

type SortedEntry struct {
	key string
	value uint64
}

type PresortedRanker struct {
	sortedVals []*SortedEntry
	sortedKeys []*SortedEntry
}

func (r *PresortedRanker) TopN(n uint64) []uint64 {
	ret := make([]uint64, n)
	for i, v := range r.sortedVals[0:int(n)] {
		ret[i] = v.value
	}
	return ret
}

func (r *PresortedRanker) Insert(key string, value uint64) {
	var index int
	if len(r.sortedVals) == 0 {
		index = 0
	} else {
		index = r.getIndex(value)
	}
	entry := &SortedEntry{key, value}
	tail := append([]*SortedEntry{entry}, r.sortedVals[index:len(r.sortedVals)]...)
	r.sortedVals = append(r.sortedVals[0:index], tail...)

	var keyIndex int
	if len(r.sortedKeys) == 0 {
		index = 0
	} else {
		index, _ = r.getEntry(key)
	}
	keyTail := append([]*SortedEntry{entry}, r.sortedKeys[keyIndex:len(r.sortedKeys)]...)
	r.sortedKeys = append(r.sortedKeys[0:keyIndex], keyTail...)
}

func (r *PresortedRanker) getIndex(value uint64) int {
	return r._getIndex(0, len(r.sortedVals) - 1, value)
}

func (r *PresortedRanker) _getIndex(start int, end int, value uint64) int {
	if end == start {
		if value < r.sortedVals[end].value {
			return end + 1
		} else {
			return end
		}
	}
	delta := end - start
	middle := start + (delta / 2)
	cur := r.sortedVals[middle]

	if cur.value == value {
		return middle
	} else if value < cur.value {
		return r._getIndex(middle + 1, end, value)
	} else {
		return r._getIndex(start, middle, value)
	}
}

func (r *PresortedRanker) getEntry(key string) (int, *SortedEntry) {
	return r._getEntry(0, len(r.sortedKeys) - 1, key)
}

func (r *PresortedRanker) _getEntry(start int, end int, value string) (int, *SortedEntry) {
	if end == start {
		if strings.Compare(value, r.sortedKeys[end].key) == 0 {
			return end, r.sortedVals[end]
		} else {
			return end, nil
		}
	}
	delta := end - start
	middle := start + (delta / 2)
	cur := r.sortedVals[middle]

	if strings.Compare(cur.key, value) == 0 {
		return middle, r.sortedVals[middle]
	} else if strings.Compare(cur.key, value) > 0 {
		return r._getEntry(middle + 1, end, value)
	} else {
		return r._getEntry(start, middle, value)
	}
}

func (r *PresortedRanker) Get(key string) uint64 {
	if _, val := r.getEntry(key); val != nil {
		return val.value
	}
	return 0
}

func (r *PresortedRanker) Delete(key string) bool {

	if keyIndex, val := r.getEntry(key); val != nil {
		valIndex := r.getIndex(val.value)
		r.sortedKeys = append(r.sortedKeys[0:keyIndex], r.sortedKeys[keyIndex:]...)
		r.sortedVals = append(r.sortedVals[0:valIndex], r.sortedVals[valIndex:]...)
		return true
	}
	return false
}
