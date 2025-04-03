package utils

type Either[L interface{}, R interface{}] interface {
	Fold(func(L), func(R))
	Match(func(L) bool, func(R) bool) bool
	IsLeft() bool
	IsRight() bool
}

type Right[L interface{}, R interface{}] struct {
	value R
}

func NewRight[L interface{}, R interface{}](value R) Right[L, R] {
	return Right[L, R]{value: value}
}

func (r Right[L, R]) Fold(_ func(L), onRight func(R)) {
	onRight(r.value)
}

func (r Right[L, R]) Match(_ func(L) bool, onRight func(R) bool) bool {
	return onRight(r.value)
}

func (r Right[L, R]) IsLeft() bool {
	return false
}

func (r Right[L, R]) IsRight() bool {
	return true
}

type Left[L interface{}, R interface{}] struct {
	value L
}

func NewLeft[L interface{}, R interface{}](value L) Left[L, R] {
	return Left[L, R]{value: value}
}

func (l Left[L, R]) Fold(onLeft func(L), _ func(R)) {
	onLeft(l.value)
}

func (l Left[L, R]) Match(onLeft func(L) bool, _ func(R) bool) bool {
	return onLeft(l.value)
}

func (l Left[L, R]) IsLeft() bool {
	return true
}

func (l Left[L, R]) IsRight() bool {
	return false
}

func FailureOf[R interface{}](err error, code int) Either[Failure, R] {
	return NewLeft[Failure, R](NewUnknownFailure(err.Error(), &code))
}
