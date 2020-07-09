package assert

import (
	"fmt"
	"testing"
)

// That ...
func That(name string, t *testing.T, got, expected interface{}) {
	t.Run(name, func(t *testing.T) {
		a := fmt.Sprintf("%v", got)
		b := fmt.Sprintf("%v", expected)
		if a != b {
			t.Errorf("got [%v] but expected [%v]", got, expected)
		}
	})
}
