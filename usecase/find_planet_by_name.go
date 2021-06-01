package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type (
	// FindPlanetByNameUseCase input port
	FindPlanetByNameUseCase interface {
		Execute(context.Context, FindPlanetByNameInput) (domain.Planet, error)
	}

	// FindPlanetByNameInput input data
	FindPlanetByNameInput struct {
		ID string
	}

	// FindPlanetByNamePresenter output port
	FindPlanetByNamePresenter interface {
		Output(domain.Planet) domain.Planet
	}

	findPlanetByNameInteractor struct {
		repository domain.PlanetFinderByName
		presenter  FindPlanetByNamePresenter
		ctxTimeout time.Duration
	}
)

// NewFindPlanetByNameInteractor creates new findPlanetByNameInteractor with its dependencies
func NewFindPlanetByNameInteractor(
	repository domain.PlanetFinderByName,
	presenter FindPlanetByNamePresenter,
	ctxTimeout time.Duration,
) FindPlanetByNameUseCase {
	return findPlanetByNameInteractor{repository: repository, presenter: presenter, ctxTimeout: ctxTimeout}
}

// Execute orchestrates the use case
func (f findPlanetByNameInteractor) Execute(ctx context.Context, input FindPlanetByNameInput) (domain.Planet, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planet, err := f.repository.FindByName(ctx, input.ID)
	if err != nil {
		return f.presenter.Output(domain.Planet{}), err
	}

	return f.presenter.Output(planet), nil
}
