package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetFinderByIDRepository struct {
	result domain.Planet
	err    error
}

func (s stubPlanetFinderByIDRepository) FindByID(_ context.Context, _ string) (domain.Planet, error) {
	return s.result, s.err
}

type stubFindPlanetByIDPresenter struct{}

func (s stubFindPlanetByIDPresenter) Output(planet domain.Planet) domain.Planet {
	return planet
}

func Test_findByIDPlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetFinderByID
		presenter  FindPlanetByIDPresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx   context.Context
		input FindPlanetByIDInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Planet
		wantErr bool
	}{
		{
			name: "Should return a planet by id",
			fields: fields{
				repository: stubPlanetFinderByIDRepository{
					result: domain.NewPlanet(
						"fakeID",
						"fakeName",
						"fakeClimate",
						"fakeTerrain",
						time.Time{},
					),
					err: nil,
				},
				presenter:  stubFindPlanetByIDPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByIDInput{
					ID: "fakeID",
				},
			},
			want: domain.NewPlanet(
				"fakeID",
				"fakeName",
				"fakeClimate",
				"fakeTerrain",
				time.Time{},
			),
			wantErr: false,
		},
		{
			name: "Should return that the planet was not found",
			fields: fields{
				repository: stubPlanetFinderByIDRepository{
					result: domain.Planet{},
					err:    domain.ErrPlanetNotFound,
				},
				presenter:  stubFindPlanetByIDPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByIDInput{
					ID: "fakeID",
				},
			},
			want:    domain.Planet{},
			wantErr: true,
		},
		{
			name: "Should fail to return a planet by id",
			fields: fields{
				repository: stubPlanetFinderByIDRepository{
					result: domain.Planet{},
					err:    errors.New("failed find planet by id"),
				},
				presenter:  stubFindPlanetByIDPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: FindPlanetByIDInput{
					ID: "fakeID",
				},
			},
			want:    domain.Planet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewFindPlanetByIDInteractor(
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
