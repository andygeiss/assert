package assert

// Matcher provides a function to match a given value.
type Matcher interface {
	Matches(interface{}) bool
	Value() interface{}
}
