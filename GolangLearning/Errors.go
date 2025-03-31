package golanglearning
import (
	"fmt"
)

type invalidDataError struct{
	message string
	field string
}

func NewInvalidDataError(message string,field string) *invalidDataError{
	return &invalidDataError{message: message, field: field}
}

func (e * invalidDataError)Error() string{
	return fmt.Sprintf("Invalid data error:%s",e.message)
}

func (e *invalidDataError) GetInvalidField() string{
	return e.field
}
