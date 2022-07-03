package shipyard

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

const (
	delimiter string = "/"
	projects  string = "projects" + delimiter
	tokens    string = "tokens" + delimiter
)

type Project struct {
	ID          uuid.UUID      `gorm:"primaryKey;default:uuid_generate_v4()" json:"-"`
	DisplayName string         `json:"displayName"`
	Description string         `json:"description"`
	Logo        string         `json:"logo"`
	TokenID     uuid.UUID      `json:"tokenId"`
	MaxSupply   float32        `json:"maxSupply"`
	InitSupply  float32        `json:"initSupply"`
	Links       datatypes.JSON `json:"links"`
	Token       token.Token    `json:"token"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (p *Project) ToProtobuf() (*api.Project, error) {
	projectPB := new(api.Project)

	projectJSON, err := json.Marshal(p)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal project: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	err = protojson.Unmarshal(projectJSON, projectPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal project: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	projectPB.Name = projects + p.ID.String()
	projectPB.Token.Name = tokens + p.Token.ID.String()

	return projectPB, nil
}

func NewProject(projectPB *api.Project) (*Project, error) {
	project := new(Project)
	projectJSON, err := protojson.Marshal(projectPB)

	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal project protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	err = json.Unmarshal(projectJSON, project)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal project protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	if projectPB.Name != "" {
		projectUUID := strings.Split(projectPB.Name, delimiter)[1]
		project.ID, err = uuid.Parse(projectUUID)
		if err != nil {
			errMsg := fmt.Sprintf("failed to parse project uuid: %v", err)
			log.Error(errMsg)
			return nil, status.Error(codes.Aborted, errMsg)
		}
	}

	return project, nil
}
