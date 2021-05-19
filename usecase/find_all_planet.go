package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type (
	// FindAllPlanetUseCase input port
	FindAllPlanetUseCase interface {
		Execute(context.Context) ([]FindAllPlanetOutput, error)
	}

	// FindAllPlanetPresenter output port
	FindAllPlanetPresenter interface {
		Output([]domain.Planet) []FindAllPlanetOutput
	}

	// FindAllPlanetOutput output data
	FindAllPlanetOutput struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Climate    string `json:"climate"`
		Ground     string `json:"ground"`
		FindAlldAt string `json:"created_at"`
	}

	findAllPlanetInteractor struct {
		repository domain.PlanetFinder
		presenter  FindAllPlanetPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllPlanetInteractor(
	repository domain.PlanetFinder,
	presenter FindAllPlanetPresenter,
	ctxTimeout time.Duration,
) FindAllPlanetUseCase {
	return findAllPlanetInteractor{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

func (f findAllPlanetInteractor) Execute(ctx context.Context) ([]FindAllPlanetOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, f.ctxTimeout*time.Second)
	defer cancel()

	planets, err := f.repository.FindAll(ctx)
	if err != nil {
		return f.presenter.Output([]domain.Planet{}), err
	}

	return f.presenter.Output(planets), nil
}
