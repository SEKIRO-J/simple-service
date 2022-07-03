package app

import (
	"context"

	api "github.com/sekiro-j/simpleservice/api/protos/v1"
	"github.com/sekiro-j/simpleservice/internal/db"
)

type Server struct {
	dbc *db.Connection
	api.UnimplementedSimpleServiceServer
}

func New(dbc *db.Connection) *Server {
	return &Server{dbc: dbc}
}

func (s *Server) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	return &api.EchoResponse{Msg: req.Msg}, nil
}
