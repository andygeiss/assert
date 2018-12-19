package is

import (
	"github.com/andygeiss/assert"
	"reflect"
)

type notEqualMatcher struct {
	val interface{}
}

// Matches ...
func (m *notEqualMatcher) Matches(val interface{}) bool {
	return !reflect.DeepEqual(m.val, val)
}

// NotEqual ...
func NotEqual(val interface{}) assert.Matcher {
	return &notEqualMatcher{val: val}
}
