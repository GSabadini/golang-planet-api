package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type (
	// CreatePlanetUseCase input port
	CreatePlanetUseCase interface {
		Execute(context.Context, CreatePlanetInput) (CreatePlanetOutput, error)
	}

	// CreatePlanetInput input data
	CreatePlanetInput struct {
		Name    string `json:"name" validate:"required"`
		Climate string `json:"climate" validate:"required"`
		Ground  int64  `json:"ground" validate:"required"`
	}

	// CreatePlanetPresenter output port
	CreatePlanetPresenter interface {
		Output(domain.Planet) CreatePlanetOutput
	}

	// CreatePlanetOutput output data
	CreatePlanetOutput struct {
		ID      string `json:"id"`
		Climate string `json:"climate"`
		Ground  int64  `json:"ground"`
	}

	createPlanetInteractor struct {
		repo       domain.PlanetCreator
		presenter  CreatePlanetPresenter
		ctxTimeout time.Duration
	}
)
