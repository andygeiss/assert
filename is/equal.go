package is

import (
	"fmt"
	"github.com/andygeiss/assert"
	"reflect"
)

type equalMatcher struct {
	val interface{}
}

// Matches ...
func (m *equalMatcher) Matches(val interface{}) bool {
	return reflect.DeepEqual(m.val, val)
}

// String ...
func (m *equalMatcher) String() string {
	return fmt.Sprintf("[%s] is equal", m.val)
}

// Equal ...
func Equal(val interface{}) assert.Matcher {
	return &equalMatcher{val: val}
}
