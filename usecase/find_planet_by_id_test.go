package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetFinder1Repository struct {
	domain.PlanetFinder
	result domain.Planet
	err    error
}

func (s stubPlanetFinder1Repository) FindByID(_ context.Context, _ string) (domain.Planet, error) {
	return s.result, s.err
}

type stubFindByIDPlanetPresenter struct{}

func (s stubFindByIDPlanetPresenter) Output(planet domain.Planet) FindByIDPlanetOutput {
	return FindByIDPlanetOutput{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Ground:  planet.Ground(),
	}
}

func Test_findByIDPlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetFinder
		presenter  FindByIDPlanetPresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx context.Context
		ID  string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    FindByIDPlanetOutput
		wantErr bool
	}{
		{
			name: "Should return a planet by id",
			fields: fields{
				repository: stubPlanetFinder1Repository{
					result: domain.NewPlanet(
						"fakeID",
						"fakeName",
						"fakeClimate",
						"fakeGround",
						time.Time{},
					),
					err: nil,
				},
				presenter:  stubFindByIDPlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				ID:  "fakeID",
			},
			want: FindByIDPlanetOutput{
				ID:      "fakeID",
				Name:    "fakeName",
				Climate: "fakeClimate",
				Ground:  "fakeGround",
			},
			wantErr: false,
		},
		{
			name: "Should return that the planet was not found",
			fields: fields{
				repository: stubPlanetFinder1Repository{
					result: domain.Planet{},
					err:    domain.ErrPlanetNotFound,
				},
				presenter:  stubFindByIDPlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				ID:  "fakeID",
			},
			want:    FindByIDPlanetOutput{},
			wantErr: true,
		},
		{
			name: "Should fail to return a planet by id",
			fields: fields{
				repository: stubPlanetFinder1Repository{
					result: domain.Planet{},
					err:    errors.New("failed find planet by id"),
				},
				presenter:  stubFindByIDPlanetPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				ID:  "fakeID",
			},
			want:    FindByIDPlanetOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewFindByIDPlanetInteractor(
				tt.fields.repository,
				tt.fields.presenter,
				tt.fields.ctxTimeout,
			)

			got, err := interactor.Execute(tt.args.ctx, tt.args.ID)
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
