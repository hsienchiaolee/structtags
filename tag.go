package structtags

import "strings"

// Tag represents a single Tag object
type Tag struct {
	Key     string
	Name    string
	Options []string
}

// Contains checks for existance of option in Tag
func (t *Tag) Contains(option string) bool {
	for _, o := range t.Options {
		if o == option {
			return true
		}
	}
	return false
}

// parseTag takes a single tag string and returns *Tag
// e.g. json:"name,omitempty" -> &Tag{Key:json Name:name Options:[omitempty]}
func parseTag(s string) *Tag {
	keyIndex := strings.Index(s, ":")
	if keyIndex == -1 {
		return nil
	}

	values := strings.Split(strings.Trim(s[keyIndex+1:], "\""), ",")
	if len(values) == 0 {
		return nil
	}

	return &Tag{
		Key:     s[:keyIndex],
		Name:    values[0],
		Options: values[1:],
	}
}
