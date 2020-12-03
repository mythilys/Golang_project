package main

import "fmt"

func perm(j []int) {
	for i := len(j) - 1; i >= 0; i-- {
		if i == 0 || j[i] < len(j)-i-1 {
			j[i]++
			return
		}
		j[i] = 0
	}
}

func nextperm(kitkat, j []int) []int {
	result := append([]int{}, kitkat...)
	for i, v := range j {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func main() {
	kitkat := []int{12, 22, 33}
	for j := make([]int, len(kitkat)); j[0] < len(j); perm(j) {
		fmt.Println(nextperm(kitkat, j))
	}
}
