package bench_util

import (
	"github.com/dinocorpllc/golang-perf"
	"math/rand"
	"github.com/satori/go.uuid"
	"testing"
)

var testKeys = make([]string, 0)
var curKey = 0

// package level storage to prevent compiler optimization
var topNResult []uint64
var getResult uint64
var keyResult string
var okResult bool

const (
	initialEntryCount = 10000
	seed = 3189080989018
)

func InitRanker(r ranker.Ranker) {
	rand.Seed(seed)
	for x := 0; x < initialEntryCount; x++ {
		key := uuid.Must(uuid.NewV4()).String()
		val := rand.Uint64()
		r.Insert(key, val)
		testKeys = append(testKeys, key)
	}
}

func BenchRankerInsert(b *testing.B, r ranker.Ranker) {
	var ok bool
	for n:=0;n<b.N;n++ {
		r.Insert(uuid.Must(uuid.NewV4()).String(), rand.Uint64())
	}
	okResult = ok
}

func BenchRankerDelete(b *testing.B, r ranker.Ranker) {
	var ok bool
	for n:=0;n<b.N;n++ {
		k := getKey()
		ok = r.Delete(k)
	}
	okResult = ok
}

func BenchRankerTopN(b *testing.B, r ranker.Ranker, n uint64) {
	var result []uint64
	for i:=0;i<b.N;i++ {
		result = r.TopN(n)
	}
	topNResult = result
}

func BenchRankerGet(b *testing.B, r ranker.Ranker) {
	var result uint64
	for i:=0; i<b.N;i++ {
		result = r.Get(getKey())
	}
	getResult = result
}

// we need to subtract the time used to get keys off of Get/Delete benchmarks
func BenchGetKey(b *testing.B, r ranker.Ranker) {
	var key string
	for i:=0;i<b.N;i++ {
		key = getKey()
	}
	keyResult = key
}

func getKey() string {
	curKey += 1
	return testKeys[curKey % len(testKeys)]
}