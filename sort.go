package sort

import (
	"math/rand"

	"github.com/fabioberger/data-structures/queue"
)

/* QuickSort
* Worst Case: O(n^2)
* Best Case: O(n*h)
* Expected Case: O(nlogn) if given randomly ordered data to sort
* if not random, since we pick pivot point at end of array, a sorted
* array will make the pivot lead only to n-1 partition!
* NOTE: For guarenteed O(nlogn) runtime, we randomly permute input in O(n) prior to running
* Benchmarks
* BenchmarkQuickSort5	 5000000	      354 ns/op
* BenchmarkQuickSort50	500000	      3874 ns/op
* BenchmarkQuickSort500	 50000	      45628 ns/op
* BenchmarkQuickSort5000  5000	      532148 ns/op
 */
func QuickSort(items []int) {
	// randomly permute inputs
	for i := range items {
		j := rand.Intn(i + 1)
		items[i], items[j] = items[j], items[i]
	}
	// call recursive QuickSort (Qs)
	Qs(items, 0, len(items)-1)
}

// Recursive Quicksort function
func Qs(items []int, low int, high int) {
	if (high - low) > 0 {
		p := qs_Partition(items, low, high)
		Qs(items, low, p-1)
		Qs(items, p+1, high)
	}
}

// Qs_Partition completes the partition step of quicksort
func qs_Partition(item []int, l int, h int) int {
	p := h
	firstHigh := l
	for i := l; i < h; i++ {
		if item[i] < item[p] {
			item[i], item[firstHigh] = item[firstHigh], item[i] //Swap
			firstHigh++
		}
	}
	item[p], item[firstHigh] = item[firstHigh], item[p] //Swap
	return firstHigh
}

/* Merge Sort
* O(nlogn)
 */
func MergeSort(items []int) {
	Ms(items, 0, len(items)-1)
}

// Ms is the recursive mergesort function
func Ms(items []int, low int, high int) {
	if low < high {
		middle := (low + high) / 2
		Ms(items, middle+1, high)
		Ms(items, low, middle)
		Merge(items, low, middle, high)
	}
}

// Merge using fixed-size arrays and pointer to last elem used
/* Benchmarks
BenchmarkMergeSort5	 2000000	     929 ns/op
BenchmarkMergeSort50	200000	     12370 ns/op
BenchmarkMergeSort500	 10000	     140173 ns/op
BenchmarkMergeSort5000	1000	     1570014 ns/op
*/
func Merge(item []int, low int, middle int, high int) {
	pos1 := 0
	a1 := make([]int, (middle - low + 1))
	j := 0
	for i := low; i <= middle; i++ {
		a1[j] = item[i]
		j++
	}

	pos2 := 0
	a2 := make([]int, (high - middle - 1 + 1))
	j = 0
	for i := middle + 1; i <= high; i++ {
		a2[j] = item[i]
		j++
	}

	i := low
	for !(pos1 > (len(a1)-1) || pos2 > (len(a2)-1)) {
		if a1[pos1] <= a2[pos2] {
			item[i] = a1[pos1]
			i++
			pos1++
		} else {
			item[i] = a2[pos2]
			i++
			pos2++
		}
	}

	for pos1 <= (len(a1) - 1) {
		item[i] = a1[pos1]
		i++
		pos1++
	}
	for pos2 <= (len(a2) - 1) {
		item[i] = a2[pos2]
		i++
		pos2++
	}

}

// Merge using queue data stucture instead of array
/* Benchmarks
* BenchmarkMergeSort5	 1000000	    1614 ns/op
* BenchmarkMergeSort50	 50000	    31278 ns/op
* BenchmarkMergeSort500	  5000	    636638 ns/op
* BenchmarkMergeSort5000	 100	    27439425 ns/op
** Currently disabled because of lower performance **
 */
func MergeUsingQueues(item []int, low int, middle int, high int) {
	q1 := queue.NewQueue(item[low])
	for i := low + 1; i <= middle; i++ {
		q1.Enqueue(item[i])
	}
	q2 := queue.NewQueue(item[middle+1])
	for i := middle + 2; i <= high; i++ {
		q2.Enqueue(item[i])
	}

	i := low
	for !(q1.Last == nil || q2.Last == nil) {
		if q1.Last.Data <= q2.Last.Data {
			item[i], _ = q1.Dequeue()
			i++
		} else {
			item[i], _ = q2.Dequeue()
			i++
		}
	}
	for q1.Last != nil {
		item[i], _ = q1.Dequeue()
		i++
	}
	for q2.Last != nil {
		item[i], _ = q2.Dequeue()
		i++
	}
}

/* Bubble Sort
* O(n^2) time
* BenchmarkBubbleSort5	 5000000	       389 ns/op
* BenchmarkBubbleSort50	  200000	      8847 ns/op
* BenchmarkBubbleSort500	    3000	    454443 ns/op
* BenchmarkBubbleSort5000	      30	  46298065 ns/op
 */
func BubbleSort(items []int) {
	notDone := true
	for notDone {
		changes := 0
		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				changes++
			}
		}
		if changes == 0 {
			notDone = false
		}
	}
}

/* Selection Sort
* O(n^2)
* BenchmarkSelectionSort5	 5000000	       390 ns/op
* BenchmarkSelectionSort50	  300000	      5494 ns/op
* BenchmarkSelectionSort500	   10000	    164691 ns/op
* BenchmarkSelectionSort5000	     100	  11686701 ns/op
 */
func SelectionSort(items []int) {
	var min, minIndex int
	for i := 0; i < len(items); i++ {
		min, minIndex = int(^uint(0)>>1), 0 // Maximum possible int value
		for j := 0; j < len(items)-i; j++ {
			offset := j + i
			if items[offset] < min {
				min = items[offset]
				minIndex = offset
			}
		}
		items[i], items[minIndex] = min, items[i]
	}
}
