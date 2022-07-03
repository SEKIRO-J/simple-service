package app

import (
	"context"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	log "github.com/sirupsen/logrus"
)

func (s *Server) CreateProject(ctx context.Context, req *api.CreateProjectRequest) (*api.Project, error) {
	log.WithFields(log.Fields{
		"request": req,
	}).Info("creating project")

	project, err := s.dbc.CreateProject(req.GetProject())
	if err != nil {
		return nil, err
	}

	projectPB, err := project.ToProtobuf()
	if err != nil {
		return nil, err
	}

	return projectPB, nil
}
