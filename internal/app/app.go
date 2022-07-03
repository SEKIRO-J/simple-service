package app

import (
	"context"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/configs"
	"github.com/sekiro-j/metapierbackend/internal/db"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	dbc            *db.Connection
	flowScannerKey string
	api.UnimplementedMetaPierServiceServer
}

func New(dbc *db.Connection) *Server {
	envConfig, err := configs.LoadEnvConfig()
	if err != nil {
		log.Fatalf("failed to load env config: %v", err)
	}

	return &Server{dbc: dbc, flowScannerKey: envConfig.FlowScannerKey}
}

func (s *Server) UpdateFEMD(ctx context.Context, req *api.UpdateFEMDRequest) (*api.FEMD, error) {
	femdPB := req.Femd
	log.Infof("Updating frontend metadata: %v", req.Femd)

	femd, err := s.dbc.CreateFEMD(femdPB)
	if err != nil {
		return nil, err
	}

	return &api.FEMD{VersionHash: femd.VersionHash, PriceFetchInterval: femd.PriceFetchInterval}, nil
}

func (s *Server) GetFEMD(ctx context.Context, req *api.GetFEMDRequest) (*api.FEMD, error) {
	femd, err := s.dbc.GetLatestFEMD()
	if err != nil {
		return nil, err
	}

	return &api.FEMD{VersionHash: femd.VersionHash, PriceFetchInterval: femd.PriceFetchInterval}, nil
}
