package main

import "golang-datastructure-algorithm/sort"
import "fmt"

func main() {
	p := []int{8, 3, 5, 7, 11, 13}
	//	sort.InsertionSort(p)
	fmt.Println("Hello, 世界")
	newP := sort.InsertionSort(p)
	fmt.Println("the new sorrted array is ", newP)
}
