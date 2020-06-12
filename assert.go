package assert

import (
	"fmt"
	"strings"
	"testing"
)

// That ...
func That(name string, t *testing.T, got, expected interface{}) {
	t.Run(name, func(t *testing.T) {
		a := fmt.Sprintf("%v", got)
		b := fmt.Sprintf("%v", expected)
		if !strings.EqualFold(a, b) {
			t.Errorf("got [%v] but expected [%v]", got, expected)
		}
	})
}
