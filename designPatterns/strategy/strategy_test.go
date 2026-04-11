package strategy

import (
	"fmt"
	"testing"
)

func Test_strategy(t *testing.T) {
	data := []int{7, 2, 4, 9, 1}
	ctx := SortContext{}

	// Ascending sort
	ctx.SetStrategy(&AscendingSort{})
	ctx.SortData(data)
	fmt.Println("Ascending:", data)

	// Descending sort
	ctx.SetStrategy(&DescendingSort{})
	ctx.SortData(data)
	fmt.Println("Descending:", data)
}