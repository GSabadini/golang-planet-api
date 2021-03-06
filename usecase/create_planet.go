package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"

	"github.com/google/uuid"
)

type (
	// CreatePlanetUseCase input port
	CreatePlanetUseCase interface {
		Execute(context.Context, CreatePlanetInput) (domain.Planet, error)
	}

	// CreatePlanetInput input data
	CreatePlanetInput struct {
		Name    string
		Climate string
		Terrain string
	}

	// CreatePlanetPresenter output port
	CreatePlanetPresenter interface {
		Output(domain.Planet) domain.Planet
	}

	createPlanetInteractor struct {
		repository domain.PlanetCreator
		presenter  CreatePlanetPresenter
		ctxTimeout time.Duration
	}
)

// NewCreatePlanetInteractor creates new createPlanetInteractor with its dependencies
func NewCreatePlanetInteractor(
	repository domain.PlanetCreator,
	presenter CreatePlanetPresenter,
	ctxTimeout time.Duration,
) CreatePlanetUseCase {
	return createPlanetInteractor{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

// Execute orchestrates the use case
func (c createPlanetInteractor) Execute(ctx context.Context, input CreatePlanetInput) (domain.Planet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout*time.Second)
	defer cancel()

	planet, err := c.repository.Create(ctx, domain.NewPlanet(
		uuid.New().String(),
		input.Name,
		input.Climate,
		input.Terrain,
		domain.NewFilms(0),
		time.Now(),
	))
	if err != nil {
		return c.presenter.Output(domain.Planet{}), err
	}

	return c.presenter.Output(planet), nil
}
