package main

import (
	"fmt"
)

func BubbleSort(list[] int)[]int {
	for i:=1; i< len(list); i++ {
		for j:=0; j < len(list)-i; j++ {
			if (list[j] > list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}


func main() {
	var bubble []int = []int{41,23,322,14,53,77,328,133}
	fmt.Println(BubbleSort(bubble))
}