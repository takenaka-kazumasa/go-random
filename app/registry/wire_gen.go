// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/ktakenaka/go-random/app/config"
	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/interface/adaptor/restclient"
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

func InitializeSignInUsecase() *usecase.SignInUsecase {
	oauth2Config := config.GetGoogleOauthConfig()
	googleRepository := restclient.NewGoogleRepository(oauth2Config)
	userRepository := mysql.NewUserRepository()
	signInUsecase := usecase.NewSignInUsecase(googleRepository, userRepository)
	return signInUsecase
}

// usecase.go:

var (
	SampleUsecaseSet = wire.NewSet(database.MySQLConnection, sampleRepositorySet, service.NewSampleService, usecase.NewSampleUsecase)

	sampleRepositorySet = wire.NewSet(mysql.NewSampleRepository, mysql.NewTransactionManager, wire.Bind(new(repository.SampleRepository), new(*mysql.SampleRepository)), wire.Bind(new(repository.TransactionManager), new(*mysql.TransactionManager)))

	SignInUsecaseSet = wire.NewSet(
		googleRepositorySet,
		userRepositorySet, usecase.NewSignInUsecase,
	)

	googleRepositorySet = wire.NewSet(config.GetGoogleOauthConfig, restclient.NewGoogleRepository, wire.Bind(new(repository.GoogleRepository), new(*restclient.GoogleRepository)))

	userRepositorySet = wire.NewSet(mysql.NewUserRepository, wire.Bind(new(repository.UserRepository), new(*mysql.UserRepository)))
)
