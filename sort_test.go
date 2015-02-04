package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	unsorted, sorted := initList()
	QuickSort(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in quick sort result")
		}
	}
}

func TestMergesort(t *testing.T) {
	unsorted, sorted := initList()
	MergeSort(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in merge sort result")
		}
	}
}

func TestBubbleSort(t *testing.T) {
	unsorted, sorted := initList()
	BubbleSort(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in bubble sort result")
		}
	}
}

func TestSelectionSort(t *testing.T) {
	unsorted, sorted := initList()
	SelectionSort(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in selection sort result")
		}
	}
}

func initList() ([]int, []int) {
	unsorted := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
	sorted := []int{2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99}
	return unsorted, sorted
}

/*
Sorting Speed Benchmarks
*/

// Benchmarking the faster mergesort (implemented /w fixed size array & ptrs)
func benchmarkMergeSort(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		items := rand.Perm(i)
		MergeSort(items)
	}
}

func BenchmarkMergeSort5(b *testing.B)    { benchmarkMergeSort(5, b) }
func BenchmarkMergeSort50(b *testing.B)   { benchmarkMergeSort(50, b) }
func BenchmarkMergeSort500(b *testing.B)  { benchmarkMergeSort(500, b) }
func BenchmarkMergeSort5000(b *testing.B) { benchmarkMergeSort(5000, b) }

// Benchmarking Quicksort
func benchmarkQuickSort(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		items := rand.Perm(i)
		QuickSort(items)
	}
}

func BenchmarkQuickSort5(b *testing.B)    { benchmarkQuickSort(5, b) }
func BenchmarkQuickSort50(b *testing.B)   { benchmarkQuickSort(50, b) }
func BenchmarkQuickSort500(b *testing.B)  { benchmarkQuickSort(500, b) }
func BenchmarkQuickSort5000(b *testing.B) { benchmarkQuickSort(5000, b) }

// Benchmarking BubbleSort
func benchmarkBubbleSort(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		items := rand.Perm(i)
		BubbleSort(items)
	}
}

func BenchmarkBubbleSort5(b *testing.B)    { benchmarkBubbleSort(5, b) }
func BenchmarkBubbleSort50(b *testing.B)   { benchmarkBubbleSort(50, b) }
func BenchmarkBubbleSort500(b *testing.B)  { benchmarkBubbleSort(500, b) }
func BenchmarkBubbleSort5000(b *testing.B) { benchmarkBubbleSort(5000, b) }

// Benchmarking SelectionSort
func benchmarkSelectionSort(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		items := rand.Perm(i)
		SelectionSort(items)
	}
}

func BenchmarkSelectionSort5(b *testing.B)    { benchmarkSelectionSort(5, b) }
func BenchmarkSelectionSort50(b *testing.B)   { benchmarkSelectionSort(50, b) }
func BenchmarkSelectionSort500(b *testing.B)  { benchmarkSelectionSort(500, b) }
func BenchmarkSelectionSort5000(b *testing.B) { benchmarkSelectionSort(5000, b) }
