package frontend

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

type Metadata struct {
	Id                 uuid.UUID      `gorm:"primaryKey; default:uuid_generate_v4()"`
	VersionHash        string         `json:"versionHash"`
	PriceFetchInterval string         `json:"priceFetchInterval"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-"`
}

func NewFEMD(femdPB *api.FEMD) (*Metadata, error) {
	femd := new(Metadata)
	femdJSON, err := protojson.Marshal(femdPB)
	if err != nil {
		errMsg := "failed to marshal frontend metadata protobuf"
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}
	err = json.Unmarshal(femdJSON, femd)
	if err != nil {
		errMsg := "failed to unmarshal frontend metadata protobuf"
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return femd, nil
}
