package datasource

type UpdateInputByID struct {
	Source string
	Data   map[string]interface{}
	ID     string
}

func NewUpdateInput[T interface{}](source string, data map[string]interface{}, ID string) *UpdateInputByID {
	return &UpdateInputByID{
		Source: source,
		Data:   data,
		ID:     ID,
	}
}
