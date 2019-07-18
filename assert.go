package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// That checks if a given value matches.
func That(t *testing.T, value interface{}, isEqual interface{}) {
	if !reflect.DeepEqual(value, isEqual) {
		msg := fmt.Sprintf("[%v] does not match [%v]!", value, isEqual)
		t.Fatal(msg)
	}
}
