package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	gSize   = 10
	gFactor = 10000
	gArr    = make([]string, gSize)
	gMap    = map[string]struct{}{}
)

func main() {
	sizes := []int{5, 10, 50, 100, 500}
	for _, size := range sizes {
		gSize = size
		gArr = make([]string, gSize)
		gMap = map[string]struct{}{}
		initData()
		arrDoneIn := checkArr()
		mapDoneIn := checkMap()
		fmt.Printf("benchmark for size %d\n", size)
		fmt.Printf("arr lookups done in %d nanosecs (%.2f lookup/sec)\nmap lookups done in %d nanosecs (%.2f lookup/sec)\nDiff (arr perforamnce is 100%%): %.2f\n", arrDoneIn, float64(arrDoneIn/int64(gSize*gFactor)),
			mapDoneIn, float64(mapDoneIn/int64(gSize*gFactor)),
			float64(mapDoneIn*100)/float64(arrDoneIn))
	}
}

func checkArr() int64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	total := int64(0)
	for i := 0; i < gSize*gFactor; i++ {
		key := fmt.Sprintf("somevaluetolookfor_%5d", rnd.Intn(gSize*gSize))
		start := time.Now().UnixNano()
		isInArray(key)
		doneIn := time.Now().UnixNano() - start
		total += doneIn
	}
	return total
}

func isInArray(key string) bool {
	for i := 0; i < gSize; i++ {
		if gArr[i] == key {
			return true
		}
	}
	return false
}

func checkMap() int64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	ok := false
	total := int64(0)
	for i := 0; i < gSize*gFactor; i++ {
		key := fmt.Sprintf("somevaluetolookfor_%5d", rnd.Intn(gSize*gSize))
		start := time.Now().UnixNano()
		_, ok = gMap[key]
		doneIn := time.Now().UnixNano() - start
		total += doneIn
	}
	if ok {
		return total
	}
	return total
}

func initData() {
	for i := 0; i < gSize; i++ {
		key := fmt.Sprintf("somevaluetolookfor_%5d", i)
		gArr[i] = key
		gMap[key] = struct{}{}
	}
}
