package structtags

import (
	"reflect"
	"strings"
)

// StructTags is a collect of Tag structs
// for all properties in a struct
type StructTags struct {
	tags []*Tags
}

// Tags returns all Tags with matching tag key
func (st *StructTags) Tags(key string) []*Tag {
	var tags []*Tag
	for _, ts := range st.tags {
		if t := ts.Get(key); t != nil {
			tags = append(tags, t)
		}
	}
	return tags
}

// Parse creates a StructTag from an object
func Parse(object interface{}) *StructTags {
	var tags []*Tags

	t := reflect.TypeOf(object)
	for i := 0; i < t.NumField(); i++ {
		if ts := parseTags(strings.Trim(string(t.Field(i).Tag), " ")); len(ts.tags) > 0 {
			tags = append(tags, ts)
		}
	}
	return &StructTags{tags: tags}
}
