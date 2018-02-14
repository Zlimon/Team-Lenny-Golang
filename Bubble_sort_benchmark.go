package algorithms

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)

func BubbleSort(list[] int)[]int {
	for i:=1; i< len(list); i++ {
		for j:=0; j < len(list)-i; j++ {
			if (list[j] > list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

// https://golang.org/doc/effective_go.html#init
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

// Skriv "benchmark"-tester for benchmarkBSortModified funksjonen
// Skriv en ny testfunksjon benchmarkBSortModified

func BenchmarkBSort100(b *testing.B) {
	benchmarkBSort(100, b)
}

func BenchmarkBSort1000(b *testing.B) {
	benchmarkBSort(1000, b)
}

func BenchmarkBSort10000(b *testing.B) {
	benchmarkBSort(10000, b)
}

func benchmarkBSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		BubbleSort(values)
	}
}

func main() {
	var bubble []int = []int{41,23,322,14,53,77,328,133,432,546,2,67,69}
	fmt.Println(BubbleSort(bubble))
}


