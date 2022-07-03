package db

import (
	"fmt"
	"strings"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (c *Connection) CreateAuction(projectName string, auctionPB *api.Auction) (*shipyard.Auction, error) {
	projectID := getID(projectName)
	project := new(shipyard.Project)
	if err := c.orm.Where("id = ?", projectID).First(project).Error; err != nil {
		errMsg := "failed to find project"
		log.WithFields(log.Fields{
			"id": projectID,
		}).Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}

	auction, err := shipyard.NewAuction(auctionPB)
	if err != nil {
		return nil, err
	}
	auction.Project = *project

	token := new(token.Token)
	if err := c.orm.Where("id = ?", auction.FundTokenID).First(token).Error; err != nil {
		errMsg := fmt.Sprintf("failed to find token: %v", err)
		log.WithFields(log.Fields{
			"id": auction.FundTokenID,
		}).Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	auction.FundToken = *token

	if err := c.orm.Save(auction).Error; err != nil {
		errMsg := fmt.Sprintf("failed to save auction: %v", err)
		log.WithFields(log.Fields{
			"auction": auction,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return auction, nil
}

func (c *Connection) ListAuctions(projectName string) ([]*shipyard.Auction, error) {
	auctions := []*shipyard.Auction{}
	err := c.orm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Scopes(filterByProjectName(tx, projectName)).Find(&auctions).Error; err != nil {
			errMsg := fmt.Sprintf("failed to list auctions: %v", err)
			log.Error(errMsg)
			return status.Error(codes.NotFound, errMsg)
		}
		return nil
	})

	return auctions, err
}

func filterByProjectName(db *gorm.DB, projectName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		projectID := getID(projectName)
		if projectID == "-" {
			return db
		}
		return db.Where("project_id = ?", projectID)
	}
}

func getID(name string) string {
	return strings.Split(name, delimiter)[1]
}
