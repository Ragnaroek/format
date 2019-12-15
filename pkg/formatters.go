package format

import (
	"fmt"
)

func applyA(arg interface{}) string {
	switch v := arg.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("<!ERROR unknown type %#v!>", v)
	}
}

func applyD(arg interface{}) string {
	return "D"
}

func applyCircumflex(arg interface{}) string {
	return "^"
}
