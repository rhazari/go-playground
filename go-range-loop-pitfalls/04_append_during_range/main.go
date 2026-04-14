package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	fmt.Println("start:", nums)
	for i, v := range nums {
		fmt.Printf("i=%d v=%d len(nums)=%d\n", i, v, len(nums))
		if v == 2 {
			nums = append(nums, 99, 100)
			fmt.Println("appended ->", nums)
		}
	}

	fmt.Println("end:", nums)
	fmt.Println("note: loop count followed the original range length")
}
