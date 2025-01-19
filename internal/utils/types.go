package utils

import "encoding/json"

type JSON = map[string]interface{}

// type Action[A interface{}] = Either[Failure, A]

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}
