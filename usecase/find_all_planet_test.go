package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetFinderAllRepository struct {
	result []domain.Planet
	err    error
}

func (s stubPlanetFinderAllRepository) FindAll(_ context.Context) ([]domain.Planet, error) {
	return s.result, s.err
}

type stubFindAllPlanetPresenter struct{}

func (s stubFindAllPlanetPresenter) Output(planets []domain.Planet) []domain.Planet {
	return planets
}

func Test_findAllPlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetFinderAll
		presenter  FindAllPlanetPresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Planet
		wantErr bool
	}{
		{
			name: "Should return a list of all the planets",
			fields: fields{
				repository: stubPlanetFinderAllRepository{
					result: []domain.Planet{
						domain.NewPlanet(
							"fakeID",
							"fakeName",
							"fakeClimate",
							"fakeTerrain",
							domain.NewFilms(0),
							time.Time{},
						),
						domain.NewPlanet(
							"fakeID2",
							"fakeName2",
							"fakeClimate2",
							"fakeTerrain2",
							domain.NewFilms(0),
							time.Time{},
						),
					},
					err: nil,
				},
				presenter:  stubFindAllPlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
			},
			want: []domain.Planet{
				domain.NewPlanet(
					"fakeID",
					"fakeName",
					"fakeClimate",
					"fakeTerrain",
					domain.NewFilms(0),
					time.Time{},
				),
				domain.NewPlanet(
					"fakeID2",
					"fakeName2",
					"fakeClimate2",
					"fakeTerrain2",
					domain.NewFilms(0),
					time.Time{},
				),
			},
			wantErr: false,
		},
		{
			name: "Should fail to return a list of all planets",
			fields: fields{
				repository: stubPlanetFinderAllRepository{
					result: []domain.Planet{},
					err:    errors.New("failed to create the planet"),
				},
				presenter:  stubFindAllPlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
			},
			want:    []domain.Planet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewFindAllPlanetInteractor(
				tt.fields.repository,
				tt.fields.presenter,
				tt.fields.ctxTimeout,
			)

			got, err := interactor.Execute(tt.args.ctx)
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
