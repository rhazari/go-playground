package strategy

import (
	"cmp"
	"slices"
)

// SortStrategy defines a common interface for sorting.
type SortStrategy interface {
	Sort(data []int)
}

// AscendingSort sorts in ascending order.
type AscendingSort struct{}

func (a *AscendingSort) Sort(data []int) {
	slices.Sort(data)
}

// DescendingSort sorts in descending order.
type DescendingSort struct{}

func (d *DescendingSort) Sort(data []int) {
	// Simple in-place reverse sort
	slices.SortFunc(data, func(a, b int) int {
		return cmp.Compare(b, a)
	})
}

// SortContext uses a SortStrategy to sort.
type SortContext struct {
	Strategy SortStrategy
}

func (c *SortContext) SetStrategy(strategy SortStrategy) {
	c.Strategy = strategy
}

func (c *SortContext) SortData(data []int) {
	if c.Strategy == nil {
		return // or panic, or default sort
	}
	c.Strategy.Sort(data)
}

