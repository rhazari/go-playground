package singleton

import (
	"fmt"
	"testing"
)

func Test_singleton(t *testing.T) {
	s1 := Instance()
	s2 := Instance()

	fmt.Println(s1 == s2)           // true: both point to the same instance
	fmt.Println(s1.Data)            // Hello, Singleton!

	s1.Data = "Changed!"
	fmt.Println(s2.Data)            // Changed! (since s1 and s2 share the same instance)
}