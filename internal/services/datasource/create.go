package datasource

type CreateInput[T interface{}] struct {
	Data   T
	Source string
}

func NewCreateInput[T interface{}](data T, source string) CreateInput[T] {
	return CreateInput[T]{
		Data:   data,
		Source: source,
	}
}
