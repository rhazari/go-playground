package main

import (
	"fmt"
	"slices"
)

func Reverse(arr []int)[]int {
	n := len(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func addToSlice(s []int)[]int {
	s = append(s, 11)
	s = append(s, 12)
	s = append(s, 13)
	return s
}

func addToSlicev2(s* []int) {
	*s = append(*s, 11)
	*s = append(*s, 12)
	*s = append(*s, 13)
}

func main() {
	fmt.Println("Slice cheat sheet")

	a := []int{12, 23, 45, 56, 76, 31, 9, 41}
	// aa := make([]int, len(a))
	// copy(aa, a)
	aa := append([]int(nil), a...)
	// Reverse
	fmt.Println("---- Reverse ----")
	slices.Reverse(aa)
	fmt.Println(aa)
	fmt.Println(a)

	b := Reverse(aa)
	fmt.Println(b)
	fmt.Println(slices.Equal(a, b))

	// Sort
	fmt.Println("---- Sorting -----")
	slices.Sort(aa)
	fmt.Println(aa)

	// value and pointer
	fmt.Println("---- Pass by value and pointer -----")
	c := []int{1, 2, 3, 4, 5}
	d := addToSlice(c)
	fmt.Println(d)

	addToSlicev2(&c)
	fmt.Println(c)

	fmt.Println("---- Append slice to another -----")
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}
	a1 = append(a1, a2...)
	fmt.Println(a1)

	fmt.Println("---- Cut slice -----")
	a1 = append(a1[:2], a1[4:]...)
	fmt.Println(a1)

	fmt.Println("---- Insert in the middle -----")
	a3 := []int{11, 12, 13}
	result := append(a1[:2], append(a3, a1[2:]...)...)
	fmt.Println(result)

	fmt.Println("---- range for loop -----")

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	a4 := [][]int{}
	for _, v := range dir {
		a4 = append(a4, v)
	}

	dir[0][0] = -10
	// both dir[i] and a4[i] point to the same underlying arrays.
	fmt.Println(dir)
	fmt.Println(a4)

}
