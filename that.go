package assert

import (
	"fmt"
	"github.com/andygeiss/log"
	"testing"
)

// That checks if state matches a given value by using a Matcher.
func That(t *testing.T, state interface{}, m Matcher) {
	if !m.Matches(state) {
		msg := fmt.Sprintf("[%v] does not match!", state)
		log.Error(msg)
	}
}
