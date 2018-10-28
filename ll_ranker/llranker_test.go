package ll_ranker

import (
	"fmt"
	"github.com/dinocorpllc/golang-perf"
	"github.com/dinocorpllc/golang-perf/bench_util"
	"testing"
)

var ll ranker.Ranker

func init() {
	ll = &LinkedListRanker{}
	bench_util.InitRanker(ll)
	fmt.Printf("Benchmarking map ranker. Top ten vals: %v\n", ll.TopN(10))
}

func BenchmarkListRankerGet(b *testing.B) {
	bench_util.BenchRankerGet(b, ll)
}

func BenchmarkGetKey(b *testing.B) {
	bench_util.BenchGetKey(b, ll)
}

func BenchmarkListRankerTop10(b *testing.B) {
	bench_util.BenchRankerTopN(b, ll, 10)
}

func BenchmarkListRankerTop100(b *testing.B) {
	bench_util.BenchRankerTopN(b, ll, 100)
}

func BenchmarkListRankerTop1000(b *testing.B) {
	bench_util.BenchRankerTopN(b, ll, 1000)
}

func BenchmarkListRankerTop10000(b *testing.B) {
	bench_util.BenchRankerTopN(b, ll, 10000)
}

func BenchmarkListRankerInsert(b *testing.B) {
	bench_util.BenchRankerInsert(b, ll)
}

func BenchmarkListRankerDelete(b *testing.B) {
	bench_util.BenchRankerDelete(b, ll)
}
