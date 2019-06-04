package structtags

import (
	"strings"
)

// Tags is a collection of Tag structs
// for a single property in a struct
type Tags struct {
	tags []*Tag
}

// Keys returns all keys of the tags
func (ts *Tags) Keys() []string {
	var keys []string
	for _, t := range ts.tags {
		keys = append(keys, t.Key)
	}
	return keys
}

// Get returns the tag matching the given key,
// returns nil if none is found
func (ts *Tags) Get(key string) *Tag {
	for _, t := range ts.tags {
		if t.Key == key {
			return t
		}
	}
	return nil
}

// parseTag creates tags using a multi-tag definition string
// e.g. db:"name" json:"name" -> *Tags{tags: []*Tag{
//   &Tag{Key:db Name:name Options:[]},
//   &Tag{Key:json Name:name Options:[]},
// }}
func parseTags(s string) *Tags {
	var tags []*Tag

	for _, t := range strings.Split(s, " ") {
		if tag := parseTag(t); tag != nil {
			tags = append(tags, tag)
		}
	}

	return &Tags{tags: tags}
}
