package db

import (
	"fmt"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/frontend"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Connection) CreateFEMD(femdPB *api.FEMD) (*frontend.Metadata, error) {
	femd, err := frontend.NewFEMD(femdPB)
	if err != nil {
		return nil, err
	}

	if err := c.orm.Create(femd).Error; err != nil {
		errMsg := fmt.Sprintf("failed to save frontend metadata: %v", err)
		log.WithFields(log.Fields{
			"femd": femd,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return femd, nil
}

func (c *Connection) GetLatestFEMD() (*frontend.Metadata, error) {
	femd := &frontend.Metadata{}
	if err := c.orm.Order("created_at desc").First(femd).Error; err != nil {
		errMsg := fmt.Sprintf("failed to save frontend metadata: %v", err)
		log.WithFields(log.Fields{
			"femd": femd,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return femd, nil
}
