package utils

func swap(arr *Intarr, i int, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

// merge two sorted arrays
func _merge(arr1 Intarr, arr2 Intarr) Intarr {
	i, j := 0, 0
	var arr3 Intarr
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		arr3 = append(arr3, arr1[i])
		i++
	}
	for j < len(arr2) {
		arr3 = append(arr3, arr2[j])
		j++
	}
	return arr3
}
