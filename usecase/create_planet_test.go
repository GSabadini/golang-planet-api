package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetCreatorRepository struct {
	result domain.Planet
	err    error
}

func (s stubPlanetCreatorRepository) Create(_ context.Context, _ domain.Planet) (domain.Planet, error) {
	return s.result, s.err
}

type stubCreatePlanetPresenter struct{}

func (s stubCreatePlanetPresenter) Output(planet domain.Planet) domain.Planet {
	return planet
}

func Test_createPlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repo       domain.PlanetCreator
		presenter  CreatePlanetPresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx   context.Context
		input CreatePlanetInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Planet
		wantErr bool
	}{
		{
			name: "Should successfully create a new planet",
			fields: fields{
				repo: stubPlanetCreatorRepository{
					result: domain.NewPlanet(
						"fakeID",
						"fakeName",
						"fakeClimate",
						"fakeTerrain",
						domain.NewFilms(0),
						time.Time{},
					),
					err: nil,
				},
				presenter:  stubCreatePlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx:   context.TODO(),
				input: CreatePlanetInput{},
			},
			want: domain.NewPlanet(
				"fakeID",
				"fakeName",
				"fakeClimate",
				"fakeTerrain",
				domain.NewFilms(0),
				time.Time{},
			),
			wantErr: false,
		},
		{
			name: "Should fail to the create a new planet",
			fields: fields{
				repo: stubPlanetCreatorRepository{
					result: domain.Planet{},
					err:    domain.ErrCreatePlanet,
				},
				presenter:  stubCreatePlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx:   context.TODO(),
				input: CreatePlanetInput{},
			},
			want:    domain.Planet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewCreatePlanetInteractor(
				tt.fields.repo,
				tt.fields.presenter,
				tt.fields.ctxTimeout,
			)

			got, err := interactor.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}
