package main

import "fmt"

type User struct {
	Name  string
	Score int
}

func main() {
	users := map[string]User{
		"u1": {Name: "Ava", Score: 10},
		"u2": {Name: "Ben", Score: 20},
	}

	for _, v := range users {
		v.Score += 5 // modifies copy only
	}
	fmt.Println("after wrong loop:", users)

	for k, v := range users {
		v.Score += 5
		users[k] = v // write back required
	}
	fmt.Println("after correct loop:", users)
}
