package structtags

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestStructTags(t *testing.T) {
	g := NewGomegaWithT(t)
	tags := Parse(struct {
		Name  string `db:"name" json:"name,omitempty"`
		Email string `db:"email" json:"email"`
		Notes string `db:"notes"`
	}{})
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
