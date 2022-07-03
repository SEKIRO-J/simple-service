package token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

const (
	delimiter string = "/"
	tokens    string = "tokens" + delimiter
)

type Token struct {
	ID          uuid.UUID      `gorm:"primaryKey;default:uuid_generate_v4()" json:"-"`
	DisplayName string         `json:"displayName"`
	Symbol      string         `json:"symbol"`
	Logo        string         `json:"logo"`
	Addr        string         `json:"addr"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func NewToken(tokenPB *api.Token) (*Token, error) {
	tokenJSON, err := protojson.Marshal(tokenPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal token protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	token := new(Token)
	err = json.Unmarshal(tokenJSON, token)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal token protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return token, nil
}

func (t *Token) ToProtobuf() (*api.Token, error) {
	tokenPB := new(api.Token)

	tokenJSON, err := json.Marshal(t)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal token: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	err = protojson.Unmarshal(tokenJSON, tokenPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal token: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	tokenPB.Name = tokens + t.ID.String()

	return tokenPB, nil
}
