// Найти 2 числа сумма которых будет равняться искомому числу 
// * массив отсортирован в порядке увеличения
// Сложность алгоритма О(n) (так как зависит от входных параметров )

package main

import (
	"fmt"
)


func twoSum(arr [7]int, sum int) []int {
	left := 0
	right := len(arr) - 1

	for left != right{
		temp := arr[left] + arr[right]
		
		if temp == sum{
			return []int{arr[left], arr[right]}
		}
		if temp < sum {
			left ++
			continue
		}
		right --
	}
	return nil 
}


func main() {
	array := [7]int{1, 3, 5, 6, 7, 8, 9}
	sum := 4
	fmt.Printf("Число %v состоиз из %v \n", sum, twoSum(array, sum))
	sum = 7
	fmt.Printf("Число %v состоиз из %v \n", sum, twoSum(array, sum))
}