package structtags

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestTags(t *testing.T) {
	g := NewGomegaWithT(t)

	tags := parseTags(`db:"name" json:"name,omitempty"`)
	g.Expect(tags.Keys()).To(Equal([]string{"db", "json"}))
	g.Expect(tags.Get("db")).To(Equal(&Tag{Key: "db", Name: "name", Options: []string{}}))
	g.Expect(tags.Get("json")).To(Equal(&Tag{Key: "json", Name: "name", Options: []string{"omitempty"}}))
	g.Expect(tags.Get("nonexist")).To(BeNil())
}

func ExampleTags() {
	tags := parseTags(`db:"name" json:"name,omitempty"`)
	for _, t := range tags.tags {
		fmt.Printf("%+v\n", t)
	}
	// Output:
	// &{Key:db Name:name Options:[]}
	// &{Key:json Name:name Options:[omitempty]}
}
