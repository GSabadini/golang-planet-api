package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetFinderByNameRepository struct {
	result domain.Planet
	err    error
}

func (s stubPlanetFinderByNameRepository) FindByName(_ context.Context, _ string) (domain.Planet, error) {
	return s.result, s.err
}

type stubFindPlanetByNamePresenter struct{}

func (s stubFindPlanetByNamePresenter) Output(planet domain.Planet) domain.Planet {
	return planet
}

func Test_findPlanetByNameInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetFinderByName
		presenter  FindPlanetByNamePresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx   context.Context
		input FindPlanetByNameInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Planet
		wantErr bool
	}{
		{
			name: "Should return a planet by name",
			fields: fields{
				repository: stubPlanetFinderByNameRepository{
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
				presenter:  stubFindPlanetByNamePresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByNameInput{
					ID: "fakeName",
				},
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
			name: "Should return that the planet was not found",
			fields: fields{
				repository: stubPlanetFinderByNameRepository{
					result: domain.Planet{},
					err:    domain.ErrPlanetNotFound,
				},
				presenter:  stubFindPlanetByNamePresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByNameInput{
					ID: "fakeName",
				},
			},
			want:    domain.Planet{},
			wantErr: true,
		},
		{
			name: "Should fail to return a planet by name",
			fields: fields{
				repository: stubPlanetFinderByNameRepository{
					result: domain.Planet{},
					err:    errors.New("failed find planet by name"),
				},
				presenter:  stubFindPlanetByNamePresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByNameInput{
					ID: "fakeName",
				},
			},
			want:    domain.Planet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewFindPlanetByNameInteractor(
				tt.fields.repository,
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
