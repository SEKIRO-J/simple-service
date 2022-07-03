package shipyard

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type Auction struct {
	ID          uuid.UUID      `gorm:"primaryKey; default:uuid_generate_v4()" json:"-"`
	ProjectID   uuid.UUID      `json:"projectId"`
	PoolID      string         `json:"poolId"`
	SaleType    string         `json:"saleType"`
	Offer       float32        `json:"offer"`
	Remaining   float32        `json:"remaining"`
	FundTokenID uuid.UUID      `json:"fundTokenId"`
	Target      float32        `json:"target"`
	Stages      datatypes.JSON `json:"stages"`
	Project     Project        `json:"-"`
	FundToken   token.Token    `json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

const (
	auctions string = "auctions" + delimiter
)

func (a *Auction) ToProtobuf() (*api.Auction, error) {
	auctionPB := new(api.Auction)

	auctionJSON, err := json.Marshal(a)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal auction: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}
	err = protojson.Unmarshal(auctionJSON, auctionPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal auction: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	auctionPB.FundToken, err = a.FundToken.ToProtobuf()
	if err != nil {
		return nil, err
	}
	auctionPB.Project, err = a.Project.ToProtobuf()
	if err != nil {
		return nil, err
	}

	auctionPB.Name = projects + a.ProjectID.String() + delimiter + auctions + a.ID.String()

	return auctionPB, nil
}

func BatchConvertAuctions(auctions []*Auction) ([]*api.Auction, error) {
	auctionPBs := make([]*api.Auction, len(auctions))
	for i, auction := range auctions {
		auctionPB, err := auction.ToProtobuf()
		if err != nil {
			return auctionPBs, err
		}
		auctionPBs[i] = auctionPB
	}

	return auctionPBs, nil
}

func NewAuction(auctionPB *api.Auction) (*Auction, error) {
	auction := new(Auction)
	auctionJSON, err := protojson.Marshal(auctionPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal auction protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}
	err = json.Unmarshal(auctionJSON, auction)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal auction protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	if auctionPB.Name != "" {
		auctionUUID := strings.Split(auctionPB.Name, delimiter)[3]
		auction.ID, err = uuid.Parse(auctionUUID)
		if err != nil {
			errMsg := fmt.Sprintf("failed to parse auction uuid: %v", err)
			log.Error(errMsg)
			return nil, status.Error(codes.Aborted, errMsg)
		}
	}

	return auction, nil
}
