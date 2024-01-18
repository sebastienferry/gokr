package customerrors

import "fmt"

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Argument)
}

type ArgumentError struct {
	Message  string
	Argument string
}

func NewArgumentError(argument string) ArgumentError {
	return ArgumentError{Message: "Argument error", Argument: argument}
}

var ErrExisting error = CustomError{Message: "Already existing"}
var ErrNotFound error = CustomError{Message: "Not found"}
