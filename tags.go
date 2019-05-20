package structtags

import (
	"strings"
)

type Tags struct {
	tags []*Tag
}

func (ts *Tags) Keys() []string {
	var keys []string
	for _, t := range ts.tags {
		keys = append(keys, t.Key)
	}
	return keys
}

func (ts *Tags) Get(key string) *Tag {
	for _, t := range ts.tags {
		if t.Key == key {
			return t
		}
	}
	return nil
}

func parseTags(s string) *Tags {
	var tags []*Tag

	for _, t := range strings.Split(s, " ") {
		if tag := parseTag(t); tag != nil {
			tags = append(tags, tag)
		}
	}

	return &Tags{tags: tags}
}
