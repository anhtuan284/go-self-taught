package algorithms

func QuickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		p := partition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
}

// partition rearranges the elements so that elements less than the pivot
// are on the left, and elements greater than the pivot are on the right.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the rightmost element as pivot
	i := low - 1       // Index of the smaller element

	for j := low; j < high; j++ {
		// If current element is less than or equal to pivot, swap it
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap the pivot element with the element at i+1
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // Return the partition index
}
