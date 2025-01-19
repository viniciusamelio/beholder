package utils

type Failure interface {
	Message() string
	Code() *int
}

type UnknownFailure struct {
	message string
	code    *int
}

func (u *UnknownFailure) Message() string {
	return u.message
}

func (u *UnknownFailure) Code() *int {
	return u.code
}

func NewUnknownFailure(message string, code *int) *UnknownFailure {
	return &UnknownFailure{message: message, code: code}
}

type NotFoundFailure struct {
	message string
	code    *int
}

func NewNotFoundFailure(message string) *NotFoundFailure {
	code := 404
	return &NotFoundFailure{message: message, code: &code}
}

func (n *NotFoundFailure) Message() string {
	return n.message
}

func (n *NotFoundFailure) Code() *int {
	return n.code
}
