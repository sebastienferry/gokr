package customerrors

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

var ErrExisting error = CustomError{Message: "Already existing"}
var ErrNotFound error = CustomError{Message: "Not found"}
