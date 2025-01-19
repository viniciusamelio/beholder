package services

import (
	"beholder-api/internal/gen/som"
	"context"
	"os"
)

var client *som.ClientImpl

type SomDatasource struct {
	Client *som.ClientImpl
}

func NewSomDatasource() *SomDatasource {
	return &SomDatasource{
		Client: client,
	}
}

func (s *SomDatasource) Environment() som.EnvironmentRepo {
	return s.Client.EnvironmentRepo()
}

func (s *SomDatasource) Session() som.SessionRepo {
	return s.Client.SessionRepo()
}

func InitSomClient() {
	somClient, err := som.NewClient(context.Background(), som.Config{
		Secure:    true,
		Host:      os.Getenv("DB_URL"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Namespace: os.Getenv("DB_NAMESPACE"),
		Database:  os.Getenv("DB_NAME"),
	})
	if err != nil {
		panic(err)
	}
	client = somClient
}
