# structtags
Go Tags helper

# Usage
```
package main

import (
	"fmt"

	"github.com/hsienchiaolee/structtags"
)

func main() {
	// Parse tags of a given struct
	tags := structtags.Parse(struct {
		Name  string `db:"name" json:"name,omitempty"`
		Email string `db:"email" json:"email"`
		Notes string `db:"notes"`
	}{})

	// Get all tags from all fields matching a struct key
	for _, t := range tags.Tags("db") {
		fmt.Printf("%+v\n", t)
	}
	// Output:
	// &{Key:db Name:name Options:[]}
	// &{Key:db Name:email Options:[]}
	// &{Key:db Name:notes Options:[]}

	// Working with a single tag
	t := tags.Tags("json")[0]
	fmt.Println(t.Key)
	fmt.Println(t.Name)
	fmt.Println(t.Options)
	fmt.Println(t.Contains("omitempty"))
	// Output:
	// json
	// name
	// [omitempty]
	// true
}
```
