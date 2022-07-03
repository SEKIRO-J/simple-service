package db

import (
	"go/token"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Connection) CreateProject(projectPB *api.Project) (*shipyard.Project, error) {

	if projectPB.TokenId == "" {
		tokenModel, err := c.CreateToken(projectPB.Token)
		if err != nil {
			return nil, err
		}
		projectPB.TokenId = tokenModel.ID.String()
	}

	// check if token is valid
	if err := c.orm.First(new(token.Token), "id = ?", projectPB.TokenId).Error; err != nil {
		errMsg := "token not found"
		log.WithFields(log.Fields{
			"token_id": projectPB.TokenId,
		}).Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}

	project, err := shipyard.NewProject(projectPB)
	if err != nil {
		return nil, err
	}

	if err := c.orm.Save(project).Error; err != nil {
		errMsg := "failed to save project"
		log.WithFields(log.Fields{
			"project": project,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return project, nil
}
