package presorted_ranker

import (
	"fmt"
	"github.com/dinocorpllc/golang-perf"
	"github.com/dinocorpllc/golang-perf/bench_util"
	"testing"
)

var r ranker.Ranker

func init() {
	r = &PresortedRanker{sortedVals: make([]*SortedEntry, 0), sortedKeys: make([]*SortedEntry, 0)}
	fmt.Printf("Benchmarking map ranker.")

	bench_util.InitRanker(r)
	fmt.Printf("Benchmarking map ranker. Top ten vals: %v\n", r.TopN(10))
}

func BenchmarkListRankerGet(b *testing.B) {
	bench_util.BenchRankerGet(b, r)
}

func BenchmarkGetKey(b *testing.B) {
	bench_util.BenchGetKey(b, r)
}

func BenchmarkListRankerTop10(b *testing.B) {
	bench_util.BenchRankerTopN(b, r, 10)
}

func BenchmarkListRankerTop100(b *testing.B) {
	bench_util.BenchRankerTopN(b, r, 100)
}

func BenchmarkListRankerTop1000(b *testing.B) {
	bench_util.BenchRankerTopN(b, r, 1000)
}

func BenchmarkListRankerTop10000(b *testing.B) {
	bench_util.BenchRankerTopN(b, r, 10000)
}

func BenchmarkListRankerInsert(b *testing.B) {
	bench_util.BenchRankerInsert(b, r)
}

func BenchmarkListRankerDelete(b *testing.B) {
	bench_util.BenchRankerDelete(b, r)
}

/**
func TestPresortedRanker_Insert(t *testing.T) {
	ranker := &PresortedRanker{sortedVals: make([]*SortedEntry, 0), index: make(map[string]*SortedEntry)}
	for x:=10; x<=15;x++{
		ranker.Insert(strconv.Itoa(x), uint64(x))
		fmt.Printf("\n vals: [")
		for _, x := range ranker.sortedVals {
			fmt.Printf("%d,", x.value)
		}
		fmt.Printf("]\n")
	}

	for x:=0; x<=5;x++{
		ranker.Insert(strconv.Itoa(x), uint64(x))
		fmt.Printf("\n vals: [")
		for _, x := range ranker.sortedVals {
			fmt.Printf("%d,", x.value)
		}
		fmt.Printf("]\n")
	}

	for x:=0; x <=16; x++{
		ranker.Insert(strconv.Itoa(x), uint64(x))
		fmt.Printf("\n vals: [")
		for _, x := range ranker.sortedVals {
			fmt.Printf("%d,", x.value)
		}
		fmt.Printf("]\n")
	}
}
*/