package usecase

import (
	"context"
	"testing"
)

func Test_uc_Tracking(t *testing.T) {
	Usecases := NewUC()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "succes",
			args:    args{ctx: context.TODO()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := Usecases.Tracking(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.Tracking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
