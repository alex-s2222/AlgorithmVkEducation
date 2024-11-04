// слияние двух отсортированных массивов

package main

import (
	"fmt"
	"time"
)

func merge(arr1 [5]int, arr2 [5]int) []int{
	merge_arr := make([]int, 0, len(arr1)+len(arr2)) 
	i, j:= 0,0 

	for i < len(arr1) && j < len(arr2){
		if arr1[i] < arr2[j]{
			merge_arr = append(merge_arr, arr1[i])
			i++
		}else{
			merge_arr = append(merge_arr, arr2[i])
			j++
		}
	}

	for i < len(arr1){
		merge_arr = append(merge_arr, arr1[i])
		i ++
	}
	for j < len(arr2) {
        merge_arr = append(merge_arr, arr2[j])
        j ++
    }
	return merge_arr
}

func optimizeMerge(arr1 [5]int, arr2 [5]int) []int{
	i, j := 0, 0
    merged := make([]int, len(arr1)+len(arr2)) // заранее выделяем нужный размер
    k := 0

    for i < len(arr1) && j < len(arr2) {
        if arr1[i] < arr2[j] {
            merged[k] = arr1[i]
            i++
        } else {
            merged[k] = arr2[j]
            j++
        }
        k++
    }

    // Копируем оставшиеся элементы
    for i < len(arr1) {
        merged[k] = arr1[i]
        i++
        k++
    }

    for j < len(arr2) {
        merged[k] = arr2[j]
        j++
        k++
    }

    return merged
}

func main(){
	arr := [5]int{1,3,5,7,9}
	arr2 := [5]int{0,2,4,6,8}

	fmt.Printf("Объединенный массив %v time %v\n", merge(arr, arr2))
	
	fmt.Printf("Объединенный массив %v\n", optimizeMerge(arr, arr2))
}