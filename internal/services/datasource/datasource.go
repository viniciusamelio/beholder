package datasource

import (
	"beholder-api/internal/utils"
)

type Datasource[T interface{}] interface {
	Create(input CreateInput[T]) utils.Either[utils.Failure, *T]
	// SelectByID(source string, ID any) utils.Either[utils.Failure, utils.JSON]
	// Select(source string, input T) utils.Either[utils.Failure, []utils.JSON]
	// Upsert(source string, input T) utils.Either[utils.Failure, utils.JSON]
	// Update(source string, input T) utils.Either[utils.Failure, utils.JSON]
	// Delete(source string, input T) utils.Either[utils.Failure, any]
}
