package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type (
	// FindAllPlanetUseCase input port
	FindAllPlanetUseCase interface {
		Execute(context.Context) ([]domain.Planet, error)
	}

	// FindAllPlanetPresenter output port
	FindAllPlanetPresenter interface {
		Output([]domain.Planet) []domain.Planet
	}

	findAllPlanetInteractor struct {
		repository domain.PlanetFinderAll
		presenter  FindAllPlanetPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllPlanetInteractor creates new findAllPlanetInteractor with its dependencies
func NewFindAllPlanetInteractor(
	repository domain.PlanetFinderAll,
	presenter FindAllPlanetPresenter,
	ctxTimeout time.Duration,
) FindAllPlanetUseCase {
	return findAllPlanetInteractor{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

// Execute orchestrates the use case
func (f findAllPlanetInteractor) Execute(ctx context.Context) ([]domain.Planet, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planets, err := f.repository.FindAll(ctx)
	if err != nil {
		return f.presenter.Output([]domain.Planet{}), err
	}

	return f.presenter.Output(planets), nil
}
