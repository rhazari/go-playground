package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	wrong := []*int{}
	for _, v := range nums {
		wrong = append(wrong, &v) // points to loop variable, not nums element
	}

	correct := []*int{}
	for i := range nums {
		correct = append(correct, &nums[i]) // points to backing array element
	}

	fmt.Println("wrong pointers:")
	for _, p := range wrong {
		fmt.Printf("%p -> %d\n", p, *p)
	}

	fmt.Println("correct pointers:")
	for _, p := range correct {
		fmt.Printf("%p -> %d\n", p, *p)
	}
}
