package usecase

import (
	"context"
	"github.com/GSabadini/golang-planet-api/domain"
	"time"
)

type (
	// FindByIDPlanetUseCase input port
	FindByIDPlanetUseCase interface {
		Execute(context.Context, string) (FindByIDPlanetOutput, error)
	}

	// FindByIDPlanetPresenter output port
	FindByIDPlanetPresenter interface {
		Output(domain.Planet) FindByIDPlanetOutput
	}

	// FindByIDPlanetOutput output data
	FindByIDPlanetOutput struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Climate   string `json:"climate"`
		Ground    string `json:"ground"`
		CreatedAt string `json:"created_at"`
	}

	findByIDPlanetInteractor struct {
		repository domain.PlanetFinder
		presenter  FindByIDPlanetPresenter
		ctxTimeout time.Duration
	}
)

// NewFindByIDPlanetInteractor creates new findByIDPlanetInteractor with its dependencies
func NewFindByIDPlanetInteractor(
	repository domain.PlanetFinder,
	presenter FindByIDPlanetPresenter,
	ctxTimeout time.Duration,
) FindByIDPlanetUseCase {
	return findByIDPlanetInteractor{repository: repository, presenter: presenter, ctxTimeout: ctxTimeout}
}

// Execute orchestrates the use case
func (f findByIDPlanetInteractor) Execute(ctx context.Context, ID string) (FindByIDPlanetOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planet, err := f.repository.FindByID(ctx, ID)
	if err != nil {
		return f.presenter.Output(domain.Planet{}), err
	}

	return f.presenter.Output(planet), nil
}
