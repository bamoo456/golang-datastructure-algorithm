package sort

// import "fmt"

func InsertionSort(arr []int) []int {
	key := 0
	for i := 1; i < len(arr); i++ {
		key = arr[i]
		z := i - 1
		// arr[z] > key for making sort asc
		// arr[z] < key for making sort desc
		for ; z >= 0 && arr[z] > key; z-- {
			arr[z+1] = arr[z]
		}
		arr[z+1] = key
	}
	return arr
}
