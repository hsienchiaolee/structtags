package structtags

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

type MyStruct struct {
	Name  string `db:"name" json:"name,omitempty"`
	Email string `db:"email" json:"email"`
	Notes string `db:"notes"`
}

func TestStructTags(t *testing.T) {
	g := NewGomegaWithT(t)
	tags := Parse(MyStruct{})
	g.Expect(tags.Tags("db")).To(Equal([]*Tag{
		&Tag{Key: "db", Name: "name", Options: []string{}},
		&Tag{Key: "db", Name: "email", Options: []string{}},
		&Tag{Key: "db", Name: "notes", Options: []string{}},
	}))
	g.Expect(tags.Tags("json")).To(Equal([]*Tag{
		&Tag{Key: "json", Name: "name", Options: []string{"omitempty"}},
		&Tag{Key: "json", Name: "email", Options: []string{}},
	}))
}

func ExampleStructTags() {
	// Parse tags of a given struct
	tags := Parse(MyStruct{})
	// Get all tags from all fields matching a struct key
	for _, t := range tags.Tags("db") {
		fmt.Printf("%+v\n", t)
	}
	// Output:
	// &{Key:db Name:name Options:[]}
	// &{Key:db Name:email Options:[]}
	// &{Key:db Name:notes Options:[]}
}
