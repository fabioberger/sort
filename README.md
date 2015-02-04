[![GoDoc](http://godoc.org/github.com/fabioberger/sort?status.svg)](https://godoc.org/github.com/fabioberger/sort)

Sorting Algorithms Implemented in Go
-----------------------------------------------

# Installation

Install with go get:

```bash
go get github.com/fabioberger/sort
```

Then include the package in your imports:

```bash
import "github.com/fabioberger/sort"
```

If you are using the sort standard library "golang.org/pkg/sort" the package names might conflict. To solve this, you can customize the name for this package:

```bash
import isort "github.com/fabioberger/sort"
```

and then reference the package with "isort" instead of "sort"

# Running Tests

To run the unit tests, execute the following from within the package directory:

```bash
go test ./... -v
```

To run the benchmark tests, execute:

```bash
go test -bench=.
```

# Example Usage

**Merge Sort:**

```go
list := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
sort.MergeSort(list)
fmt.Println(list) // [2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99]
```

**Quick Sort:**

```go
list := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
sort.QuickSort(list)
fmt.Println(list) // [2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99]
```

**Bubble Sort:**

```go
list := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
sort.BubbleSort(list)
fmt.Println(list) // [2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99]
```

**Selection Sort:**

```go
list := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
sort.SelectionSort(list)
fmt.Println(list) // [2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99]
```

# Benchmarks

The following benchmarks were achieved while running on a 2012 Macbook pro '15 with an i7 processor. It is probably most useful to observe the differences between the performances of the different sorting algorithms rather then the actual values.

**Merge Sort:**

```
BenchmarkMergeSort5	 		1000000	      	1018 ns/op
BenchmarkMergeSort50	   	100000	     	15518 ns/op
BenchmarkMergeSort500	   	10000	    	171488 ns/op
BenchmarkMergeSort5000	    1000	   		1811139 ns/op
```
Each line above represents running a single benchmarking test with the size of the list being sorted appended at the end of the first segment (i.e n = 5, 50, 500, 5000). The second column shows how many times the sort function was run (so as to get statistically significant results). As expected, the larger the input, the less times we are able to sort within the same allotted time. The last column is the average speed at which the sorting algorithm ran. Intuitively, larger inputs took longer to sort. What is very noticable is how certain sorting algorithms grow exponentially with the size of their input while others grow slower.

**Quick Sort:**

```
BenchmarkQuickSort5	 		2000000	      	 604 ns/op
BenchmarkQuickSort50	  	200000	      	 6480 ns/op
BenchmarkQuickSort500	   	20000	     	 72003 ns/op
BenchmarkQuickSort5000	    2000	    	 738737 ns/op
```
Quicksort is visibly faster on all input sizes then merge sort.

**Bubble Sort:**

```
BenchmarkBubbleSort5	 	5000000	       	390 ns/op
BenchmarkBubbleSort50	  	200000	      	8808 ns/op
BenchmarkBubbleSort500	    3000	    	455174 ns/op
BenchmarkBubbleSort5000	    30	  			46624990 ns/op
```

Bubble sort is surprisingly faster on very small input sizes but then grows exponentially with the input size to be very unusable for large problems.

**Selection Sort:**

```
BenchmarkSelectionSort5	 		5000000	       	383 ns/op
BenchmarkSelectionSort50	  	300000	      	5582 ns/op
BenchmarkSelectionSort500	   	10000	    	164446 ns/op
BenchmarkSelectionSort5000	    100	  			11579992 ns/op
```

Selection sort also grows exponentially with the size of the input but does perform better then Bubblesort if you have to pick an O(n^2) algorithm ;)