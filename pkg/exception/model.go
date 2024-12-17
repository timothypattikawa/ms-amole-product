package exception

type InternalServerError struct {
	Message string
}

func NewInternalServerError(Message string) *InternalServerError {
	return &InternalServerError{Message}
}

func (i *InternalServerError) Error() string {
	return i.Message
}

type NotFoundError struct {
	Message string
}

func NewNotFoundError(Message string) *NotFoundError {
	return &NotFoundError{Message}
}

func (i *NotFoundError) Error() string {
	return i.Message
}

type BadRequestError struct {
	Message string
}

func NewBadRequestError(Message string) *BadRequestError {
	return &BadRequestError{Message}
}

func (i *BadRequestError) Error() string {
	return i.Message
}
