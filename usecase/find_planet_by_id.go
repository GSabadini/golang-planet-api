package usecase

import (
	"context"
	"github.com/GSabadini/golang-planet-api/domain"
	"time"
)

type (
	// FindPlanetByIDUseCase input port
	FindPlanetByIDUseCase interface {
		Execute(context.Context, FindPlanetByIDInput) (domain.Planet, error)
	}

	// FindPlanetByIDInput input data
	FindPlanetByIDInput struct {
		ID string
	}

	// FindPlanetByIDPresenter output port
	FindPlanetByIDPresenter interface {
		Output(domain.Planet) domain.Planet
	}

	findPlanetByIDInteractor struct {
		repository domain.PlanetFinderByID
		presenter  FindPlanetByIDPresenter
		ctxTimeout time.Duration
	}
)

// NewFindPlanetByIDInteractor creates new findPlanetByIDInteractor with its dependencies
func NewFindPlanetByIDInteractor(
	repository domain.PlanetFinderByID,
	presenter FindPlanetByIDPresenter,
	ctxTimeout time.Duration,
) FindPlanetByIDUseCase {
	return findPlanetByIDInteractor{repository: repository, presenter: presenter, ctxTimeout: ctxTimeout}
}

// Execute orchestrates the use case
func (f findPlanetByIDInteractor) Execute(ctx context.Context, input FindPlanetByIDInput) (domain.Planet, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planet, err := f.repository.FindByID(ctx, input.ID)
	if err != nil {
		return f.presenter.Output(domain.Planet{}), err
	}

	return f.presenter.Output(planet), nil
}
