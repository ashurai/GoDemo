package main

import "fmt"

func main() {
	list := []int{-2, -3, 4, -1, -2, 1, 5, -3}

	result := SOMCA(list)
	fmt.Println("Sum: ", result[0], "i :", result[1], "j :", result[2])
}

// SOMCA Sum Of Max Contigues Array
func SOMCA(list []int) (result [3]int) {
	result = [3]int{}
	var temp int
	for i := range list {
		temp = list[i]
		for j := i + 1; j < len(list); j++ {
			temp = temp + list[j]
			if result[0] < temp {
				result[0] = temp
				result[1] = i
				result[2] = j
			}
		}
	}
	return
}

func SOMCA_Kadane(list []int) (res []int) {
	max_end := list[0]

	temp_max := list[0]

	start := 0
	end := 0
	for i := range list {
		max_end = max(list[i], max_end+list[i])
		if max_end == list[i] {
			start = i
		}
		temp_max = max(temp_max, max_end)
		if temp_max == max_end {
			end = i
		}
	}
	return []int{temp_max, start, end}
}

func max(first int, sec int) int {
	if first > sec {
		return first
	} else {
		return sec
	}
}
