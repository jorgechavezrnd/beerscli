package errors

import "fmt"

// BadResponseErr represent a response error
type BadResponseErr struct {
	Msg  string
	File string
	Line int
}

func (e *BadResponseErr) Error() string {
	return fmt.Sprintf("%s: %d: %s", e.File, e.Line, e.Msg)
}
