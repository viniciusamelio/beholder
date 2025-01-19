package services

import (
	"beholder-api/internal/services/datasource"
	"beholder-api/internal/utils"
	"fmt"
	"os"

	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)

var surreal *surrealdb.DB

func InitDB() *surrealdb.DB {
	dbUrl := os.Getenv("DB_URL")

	db, err := surrealdb.New(dbUrl)
	if err != nil {
		panic(err)
	}

	if err = db.Use(os.Getenv("DB_NAMESPACE"), os.Getenv("DB_NAME")); err != nil {
		panic(err)
	}

	authData := &surrealdb.Auth{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	_, err = db.SignIn(authData)
	if err != nil {
		panic(err)
	}

	surreal = db
	return db
}

func GetDB() *surrealdb.DB {
	return surreal
}

type SurrealDatasource[T interface{}] struct {
	db *surrealdb.DB
}

func NewSurrealDatasource[T interface{}]() SurrealDatasource[T] {
	return SurrealDatasource[T]{db: surreal}
}

func (s SurrealDatasource[T]) Create(input datasource.CreateInput[T]) utils.Either[utils.Failure, *T] {
	parsed, err := utils.StructToMap(input.Data)
	if err != nil {
		return utils.NewLeft[utils.Failure, *T](utils.NewUnknownFailure(err.Error(), nil))
	}
	parsed["id"] = utils.GenSnowflakeID()
	created, err := surrealdb.Create[T](surreal, models.Table(input.Source), parsed)
	if err != nil {
		return utils.NewLeft[utils.Failure, *T](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure, *T](created)
}

func (s SurrealDatasource[T]) UpdateByID(input datasource.UpdateInputByID) utils.Either[utils.Failure, *T] {
	queryString := "UPDATE $id SET"
	for k := range input.Data {
		queryString += fmt.Sprintf("%s: $%s,", k, k)
	}
	input.Data["id"] = input.ID

	updated, err := surrealdb.Update[T](surreal, input.ID, input.Data)
	if err != nil {
		return utils.NewLeft[utils.Failure, *T](utils.NewUnknownFailure(err.Error(), nil))
	}

	return utils.NewRight[utils.Failure, *T](updated)
}
