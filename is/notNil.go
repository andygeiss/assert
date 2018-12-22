package is

import (
	"github.com/andygeiss/assert"
	"reflect"
)
// notNilMatcher with its value will not be exported.
type notNilMatcher struct {}

// Matches checks if values are notNil.
func (m *notNilMatcher) Matches(val interface{}) bool {
	return !reflect.DeepEqual(nil, val)
}

// Value ...
func (m *notNilMatcher) Value() interface{} {
	return "not nil"
}

// NotNil initializes the notNil-Matcher.
func NotNil() assert.Matcher {
	return &notNilMatcher{}
}
