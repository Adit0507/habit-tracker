package habit

import "fmt"

type InvalidInputError struct {
	field  string
	reason string
}

func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input in field %s: %s", e.field, e.reason)
}