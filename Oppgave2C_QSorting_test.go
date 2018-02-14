package main

import (
	"testing"
	"reflect"
	"fmt"
	"time"
	"math/rand"
)

func QuickSort(b *testing.B) {
	values := []int{9, 1, 20, 3, 6, 7}
	expected := []int{1, 3, 6, 7, 9, 20}
	QuickSort(values)
	if !reflect.DeepEqual(values, expected) {
		b.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

func BenchmarkQuickSort100(b *testing.B) {
	benchmarkQSort(100, b)
}

func BenchmarkQuickSort1000(b *testing.B) {
	benchmarkQSort(1000, b)
}

func BenchmarkQuicSort10000(b *testing.B) {
	benchmarkQSort(10000, b)
}

func benchmarkQuickSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		QuickSort(values)
	}
}

func main() {
	var quick []int = []int{41,23,322,14,53,77,328,133,432,546,2,67,69}
	fmt.Println(QuickSort(quick))
}

