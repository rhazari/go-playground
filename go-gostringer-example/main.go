package main

import "fmt"

type APIKey struct {
	Owner string
	Key   string
}

// String is used by %v, %s, and fmt.Println.
func (a APIKey) String() string {
	return fmt.Sprintf("APIKey(owner=%s)", a.Owner)
}

// GoString is used by %#v when the value implements fmt.GoStringer.
func (a APIKey) GoString() string {
	return fmt.Sprintf("APIKey{Owner:%q, Key:%q}", a.Owner, "<redacted>")
}

func main() {
	k := APIKey{Owner: "backend-service", Key: "sk-live-123456"}

	fmt.Printf("%%v   => %v\n", k)
	fmt.Printf("%%#v  => %#v\n", k)
}
