package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func arraySort(xd []int) []int {
	n := len(xd)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if xd[j] < xd[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the current element
		xd[i], xd[minIndex] = xd[minIndex], xd[i]
	}
	return xd
}

func reverseArray(xd []int) []int {
	n := len(xd)
	reversedArray := make([]int, n)
	for i := 0; i < n; i++ {
		reversedArray[i] = xd[n-1]
		n -= 1
	}
	return reversedArray
}

func fillArray(length int) []int {
	var xd []int
	for i := 0; i <= length; i++ {
		xd = append(xd, rand.Intn(100))
	}
	return xd
}

func main() {
	xd := fillArray(15)
	fmt.Println("Array: ", xd)

	sortedXd := arraySort(xd)
	fmt.Println("Sorted Array: ", sortedXd)

	sort.Ints(xd)
	fmt.Println("Sorted Array Builtin Function: ", xd)

	reversedArray := reverseArray(sortedXd)
	fmt.Println("Reverse Sorted Array: ", reversedArray)

}
