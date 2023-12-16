// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package twininstance

import (
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance/controller"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance/domain/repository"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/context/twininstance/usecase"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/db"
	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/infra/validator"
)

// Injectors from wire.go:

func InitializeTwinInstanceContainer(dbConnection db.DBConnection) TwinInstanceContainer {
	twinInstanceMapper := repository.NewTwinInstanceMapper()
	twinInstanceRepository := repository.NewTwinInstanceRepository(dbConnection, twinInstanceMapper)
	twinInstanceUseCase := usecase.NewTwinInstanceUseCase(twinInstanceRepository)
	validatorValidator := validator.NewValidator()
	twinInstanceController := controller.NewTwinInstanceController(twinInstanceUseCase, validatorValidator)
	twinInstanceContainer := NewTwinInstanceContainer(twinInstanceController, twinInstanceRepository, twinInstanceUseCase)
	return twinInstanceContainer
}

// wire.go:

func NewTwinInstanceContainer(controller2 controller.TwinInstanceController, repository2 repository.TwinInstanceRepository,

	useCase usecase.TwinInstanceUseCase,
) TwinInstanceContainer {
	return TwinInstanceContainer{
		Controller: controller2,
		Repository: repository2,
		UseCase:    useCase,
	}
}

type TwinInstanceContainer struct {
	Controller controller.TwinInstanceController
	Repository repository.TwinInstanceRepository
	UseCase    usecase.TwinInstanceUseCase
}
