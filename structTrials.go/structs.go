package main

import (
	"cmp"
	"fmt"
	"slices"
)

type person struct {
	name string
	age int
	country string
	tags []int
}

func addPersonViaPointer(p person, ppl* []person) {
	(*ppl) = append((*ppl), p)
}

func addPerson(p person, ppl []person) []person {
	ppl = append(ppl, p)
	return ppl
}

func deepCopy(src person) person {
	dst := src

	if src.tags != nil {
		dst.tags = make([]int, len(src.tags))
		copy(dst.tags, src.tags)
	}
	return dst
}

func main() {
	ppl := []person{
		{
			name: "harry potter",
			age: 38,
			country: "us",
		},
		{
			name: "emily wright",
			age: 23,
			country: "us",
		},
		{
			name: "zoe parks",
			age: 47,
			country: "us",
		},
		{
			name: "joe harlow",
			age: 21,
			country: "uk",
		},
		{
			name: "jack fallow",
			age: 21,
			country: "ie",
		},
		{
			name: "sanjay sharma",
			age: 33,
			country: "in",
		},
	}
	fmt.Println(ppl)

	p1 := person{
		name: "jacky finerman",
		age: 28,
		country: "nl",
		tags: []int{2, 3, 4},
	}

	addPersonViaPointer(p1, &ppl)
	fmt.Println(ppl)

	p2 := person{
		name: "karl urban",
		age: 25,
		country: "de",
		tags: []int{7, 8, 9},
	}
	ppl = addPerson(p2, ppl)
	fmt.Println(ppl)


	// sort in ascending order of age, then country code
	slices.SortFunc(ppl, func (a, b person) int {
		if a.age == b.age {
			return cmp.Compare(a.country, b.country)
		}
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(ppl)

	// sort in descending order of age, then ascending order of country code
	slices.SortFunc(ppl, func (a, b person) int {
		if a.age == b.age {
			return cmp.Compare(a.country, b.country)
		}
		return cmp.Compare(b.age, a.age)
	})
	fmt.Println(ppl)


	// Shallow copy of slice of structs
	// newPplSlice := ppl
	// newPplSlice[0].name = "raihan hazarika"
	// fmt.Println(newPplSlice)
	// fmt.Println(ppl)

	// Deep copy of slice of structs
	newPplSlice := make([]person, len(ppl))
	copy(newPplSlice, ppl)

	for i, p := range ppl {
		newPplSlice[i] = deepCopy(p)
	}
	newPplSlice[0].name = "raihan hazarika"
	newPplSlice[0].tags = []int{11, 12, 13}
	fmt.Println(newPplSlice)
	fmt.Println(ppl)

}
