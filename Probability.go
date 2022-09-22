package main

import "fmt"

var arr = []int{3, 6, 9, 1, 0, 2, 74, 1}
var target int = 5

func twoSum(arr []int, target int) (int, int) {
	if len(arr) <= 1 {
		return 0, 0
	}
	m := make(map[int]int)
	for index, value := range arr {
		if j, ok := m[value]; ok {
			return j, index
		}
		m[target-value] = index
	}
	return 0, 0
}
func main() {
	fmt.Println(twoSum(arr, target))
}
