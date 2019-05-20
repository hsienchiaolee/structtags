package structtags

import (
	"reflect"
	"strings"
)

type StructTags struct {
	tags []*Tags
}

func (st *StructTags) Tags(key string) []*Tag {
	var tags []*Tag
	for _, ts := range st.tags {
		if t := ts.Get(key); t != nil {
			tags = append(tags, t)
		}
	}
	return tags
}

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
