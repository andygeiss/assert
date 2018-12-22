package is

import (
	"github.com/andygeiss/assert"
	"reflect"
)
// nilMatcher with its value will not be exported.
type nilMatcher struct {}

// Matches checks if values are nil.
func (m *nilMatcher) Matches(val interface{}) bool {
	return reflect.DeepEqual(nil, val)
}

// Value ...
func (m *nilMatcher) Value() interface{} {
	return nil
}

// Nil initializes the nil-Matcher.
func Nil() assert.Matcher {
	return &nilMatcher{}
}
