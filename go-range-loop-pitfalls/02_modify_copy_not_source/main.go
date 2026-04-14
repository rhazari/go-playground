package main

import "fmt"

type Item struct {
	Count int
}

func main() {
	items := []Item{{Count: 1}, {Count: 2}, {Count: 3}}

	for _, v := range items {
		v.Count += 10 // modifies copy only
	}
	fmt.Println("after wrong loop:", items)

	for i := range items {
		items[i].Count += 10 // modifies original element
	}
	fmt.Println("after correct loop:", items)
}
