package utils

import "fmt"

type Sort interface {
	BubbleSort()
	MergeSort()
}
type Intarr []int

func (arr *Intarr) BubbleSort() {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[i] {
				swap(arr, i, j)
			}
		}
	}
}

func (arr *Intarr) MergeSort() {
	//Base case
	if len((*arr)) <= 1 {
		return
	}

	arr1 := (*arr)[0 : len((*arr))/2]
	arr2 := (*arr)[len((*arr))/2 : len((*arr))]
	arr1.MergeSort()
	arr2.MergeSort()

	(*arr) = _merge(arr1, arr2)
}
func Printarr(arr Intarr) {
	fmt.Println("Printing array")
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}
