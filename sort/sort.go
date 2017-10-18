package sort

type sortFn func([]int)

func swap(arr []int, idx1, idx2 int) {
	t := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = t
}

func sort(arr []int, fn sortFn) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	fn(sorted)
	return sorted
}

// Bubblesort

// Bubblesort : Bubblesort sort a copy and return it
func Bubblesort(arr []int) []int {
	return sort(arr, BubblesortX)
}

// BubblesortX : Bubblesort sort passed array
func BubblesortX(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		swapped := false

		for i := j; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
}

// Bubblesort end

// Quicksort

func Quicksort(arr []int) []int {
	return sort(arr, QuicksortX)
}

func QuicksortX(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, low, high int) {
	if high <= low {
		return
	}
	p := partition(arr, low, high)
	quicksort(arr, low, p-1)
	quicksort(arr, p, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low+(high-low)/2]
	i := low
	j := high
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			swap(arr, i, j)
			i++
			j--
		}
	}
	return i
}
