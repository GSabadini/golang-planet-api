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
		Ground  string `json:"ground" validate:"required"`
	}

	// CreatePlanetPresenter output port
	CreatePlanetPresenter interface {
		Output(domain.Planet) CreatePlanetOutput
	}

	// CreatePlanetOutput output data
	CreatePlanetOutput struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Climate   string `json:"climate"`
		Ground    string `json:"ground"`
		CreatedAt string `json:"created_at"`
	}

	createPlanetInteractor struct {
		repo       domain.PlanetCreator
		presenter  CreatePlanetPresenter
		ctxTimeout time.Duration
	}
)

func NewCreatePlanetInteractor(repo domain.PlanetCreator, presenter CreatePlanetPresenter, ctxTimeout time.Duration) CreatePlanetUseCase {
	return createPlanetInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

func (c createPlanetInteractor) Execute(ctx context.Context, input CreatePlanetInput) (CreatePlanetOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	planet, err := c.repo.Create(ctx, domain.NewPlanet(
		"uuid.New()",
		input.Name,
		input.Climate,
		input.Ground,
		time.Now(),
	))
	if err != nil {
		return c.presenter.Output(domain.Planet{}), err
	}

	return c.presenter.Output(planet), nil
}
