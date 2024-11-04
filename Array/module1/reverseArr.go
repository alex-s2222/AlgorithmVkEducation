// Развернуть массив за линейное время

package main

import "fmt"

func reverse(arr [5]int) [5]int{
	left, right := 0, len(arr) - 1

	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left ++
		right --
	}
	return arr
}


func main(){
	arr := [5]int{1,2,3,4,5}
	fmt.Printf("%v reverse to %v \n", arr, reverse(arr))
}