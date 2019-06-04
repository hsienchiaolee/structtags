package structtags

import (
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
