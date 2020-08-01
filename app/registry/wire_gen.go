// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/app/usecase"
)

// Injectors from usecase.go:

func InitializeSampleUsecase() *usecase.SampleUsecase {
	db := database.MySQLConnection()
	sampleRepository := mysql.NewSampleRepository(db)
	transactionManager := mysql.NewTransactionManager(db)
	sampleService := service.NewSampleService(sampleRepository)
	sampleUsecase := usecase.NewSampleUsecase(sampleRepository, transactionManager, sampleService)
	return sampleUsecase
}

// usecase.go:

var (
	SampleUsecaseSet = wire.NewSet(database.MySQLConnection, sampleRepositorySet, service.NewSampleService, usecase.NewSampleUsecase)
)

var (
	sampleRepositorySet = wire.NewSet(mysql.NewSampleRepository, mysql.NewTransactionManager, wire.Bind(new(repository.SampleRepository), new(*mysql.SampleRepository)), wire.Bind(new(repository.TransactionManager), new(*mysql.TransactionManager)))
)
