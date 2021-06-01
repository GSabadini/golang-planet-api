package usecase

import (
	"context"
	"github.com/GSabadini/golang-planet-api/domain"
	"time"
)

type (
	// FindPlanetByIDUseCase input port
	FindPlanetByIDUseCase interface {
		Execute(context.Context, string) (domain.Planet, error)
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
func (f findPlanetByIDInteractor) Execute(ctx context.Context, ID string) (domain.Planet, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planet, err := f.repository.FindByID(ctx, ID)
	if err != nil {
		return f.presenter.Output(domain.Planet{}), err
	}

	return f.presenter.Output(planet), nil
}
