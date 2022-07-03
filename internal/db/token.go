package db

import (
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Connection) CreateToken(tokenPB *api.Token) (*token.Token, error) {

	token, err := token.NewToken(tokenPB)
	if err != nil {
		return nil, err
	}

	if err := c.orm.Save(token).Error; err != nil {
		errMsg := "failed to save token"
		log.WithFields(log.Fields{
			"token": token,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return token, nil
}
