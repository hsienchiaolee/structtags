package structtags

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestTag(t *testing.T) {
	g := NewGomegaWithT(t)

	tag := parseTag(`json:"name,omitempty"`)
	t.Run("parseTag", func(t *testing.T) {
		g.Expect(tag.Key).To(Equal("json"))
		g.Expect(tag.Name).To(Equal("name"))
		g.Expect(tag.Options).To(Equal([]string{"omitempty"}))
	})

	t.Run("Contains", func(t *testing.T) {
		g.Expect(tag.Contains("omitempty")).To(BeTrue())
		g.Expect(tag.Contains("notExist")).To(BeFalse())
	})
}

func ExampleTag() {
	// Parse a single tag from tag string
	tag := parseTag(`json:"name,omitempty"`)
	fmt.Println(tag.Key)
	fmt.Println(tag.Name)
	fmt.Println(tag.Options)
	fmt.Println(tag.Contains("omitempty"))
	// Output:
	// json
	// name
	// [omitempty]
	// true
}
