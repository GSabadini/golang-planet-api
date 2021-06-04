package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetDeleterRepository struct {
	err error
}

func (s stubPlanetDeleterRepository) Delete(_ context.Context, _ string) error {
	return s.err
}

func Test_deletePlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetDeleter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx   context.Context
		input DeletePlanetInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should successfully delete a planet",
			fields: fields{
				repository: stubPlanetDeleterRepository{
					err: nil,
				},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: DeletePlanetInput{
					ID: "fakeID",
				},
			},
			wantErr: false,
		},
		{
			name: "Should fail to delete a planet",
			fields: fields{
				repository: stubPlanetDeleterRepository{
					err: domain.ErrDeletePlanet,
				},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
				input: DeletePlanetInput{
					ID: "fakeID",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := NewDeletePlanetInteractor(
				tt.fields.repository,
				tt.fields.ctxTimeout,
			)

			if err := interactor.Execute(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
			}
		})
	}
}
