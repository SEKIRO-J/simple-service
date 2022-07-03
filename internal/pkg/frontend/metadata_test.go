package frontend

import (
	"reflect"
	"testing"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
)

const (
	VersionHash        = "jajajaja"
	PriceFetchInterval = "10"
)

func TestNewFEMD(t *testing.T) {
	type args struct {
		femdPB *api.FEMD
	}
	tests := []struct {
		name    string
		args    args
		want    *Metadata
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"happy path",
			args{
				&api.FEMD{
					VersionHash:        VersionHash,
					PriceFetchInterval: PriceFetchInterval,
				},
			},
			&Metadata{
				VersionHash:        VersionHash,
				PriceFetchInterval: PriceFetchInterval,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFEMD(tt.args.femdPB)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFEMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFEMD() = %v, want %v", got, tt.want)
			}
		})
	}
}
