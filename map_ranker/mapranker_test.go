package map_ranker

import (
	"fmt"
	"github.com/dinocorpllc/golang-perf"
	"github.com/dinocorpllc/golang-perf/bench_util"
	"strconv"
	"testing"
)


var mr ranker.Ranker

func init() {
	mr = &MapRanker{make(map[string]uint64)}
	bench_util.InitRanker(mr)
	fmt.Printf("Benchmarking map ranker. Top ten vals: %v\n", mr.TopN(10))
	m = make(map[string]int)
	for x := 0; x < 1000000; x++ {
		m[strconv.Itoa(x)] = x
	}
}

func BenchmarkMapRankerGet(b *testing.B) {
	bench_util.BenchRankerGet(b, mr)
}

func BenchmarkGetKey(b *testing.B) {
	bench_util.BenchGetKey(b, mr)
}

func BenchmarkMapRankerTop10(b *testing.B) {
	bench_util.BenchRankerTopN(b, mr, 10)
}

func BenchmarkMapRankerTop100(b *testing.B) {
	bench_util.BenchRankerTopN(b, mr, 100)
}

func BenchmarkMapRankerTop1000(b *testing.B) {
	bench_util.BenchRankerTopN(b, mr, 1000)
}

func BenchmarkMapRankerTop10000(b *testing.B) {
	bench_util.BenchRankerTopN(b, mr, 10000)
}

func BenchmarkMapRankerInsert(b *testing.B) {
	bench_util.BenchRankerInsert(b, mr)
}

func BenchmarkMapRankerDelete(b *testing.B) {
	bench_util.BenchRankerDelete(b, mr)
}

var m map[string]int
func BenchmarkVanillaMap(b *testing.B) {
	for x := 0; x < b.N; x++ {
		delete(m, strconv.Itoa(x))
	}
}


