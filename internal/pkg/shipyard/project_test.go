package shipyard

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	Links = map[string]interface{}{
		Discord:   Discord,
		Instagram: Instagram,
		Twitter:   Twitter,
	}
	LinksPB, _   = structpb.NewStruct(Links)
	LinksJSON, _ = json.Marshal(Links)

	ProjectID    = uuid.New()
	ProjectIDstr = ProjectID.String()

	ProjectModel = Project{
		ID:          ProjectID,
		DisplayName: DisplayName,
		Description: Description,
		Logo:        Logo,
		TokenID:     TokenID,
		MaxSupply:   MaxSupply,
		InitSupply:  InitSupply,
		Links:       LinksJSON,
		Token:       TokenModel,
	}

	ProjectPB = api.Project{
		Name:        projects + ProjectIDstr,
		DisplayName: DisplayName,
		Description: Description,
		Logo:        Logo,
		TokenId:     TokenIDstr,
		MaxSupply:   MaxSupply,
		InitSupply:  InitSupply,
		Links:       LinksPB,
		Token:       &TokenPB,
	}
)

const (
	DisplayName string  = "STEPN"
	Description string  = "The move-2-earn web3 project"
	Logo        string  = "http://www.cdn.com/project"
	MaxSupply   float32 = 12345678.12345678
	InitSupply  float32 = 678.12345678
	Discord     string  = "discord"
	Twitter     string  = "twitter"
	Instagram   string  = "instagram"
)

func Test_NewProject(t *testing.T) {
	type args struct {
		projectPB *api.Project
	}
	tests := []struct {
		name    string
		args    args
		want    *Project
		wantErr bool
	}{
		// Add test cases.
		{
			"happy path",
			args{
				&api.Project{
					Name:        "",
					DisplayName: DisplayName,
					Description: Description,
					Logo:        Logo,
					TokenId:     TokenIDstr,
					MaxSupply:   MaxSupply,
					InitSupply:  InitSupply,
				},
			},
			&Project{
				DisplayName: DisplayName,
				Description: Description,
				Logo:        Logo,
				TokenID:     TokenID,
				MaxSupply:   MaxSupply,
				InitSupply:  InitSupply,
			},
			false,
		},
		{
			"happy path with name",
			args{
				&api.Project{
					Name:        projects + ProjectIDstr,
					DisplayName: DisplayName,
					Description: Description,
					Logo:        Logo,
					TokenId:     TokenIDstr,
					MaxSupply:   MaxSupply,
					InitSupply:  InitSupply,
				},
			},
			&Project{
				ID:          ProjectID,
				DisplayName: DisplayName,
				Description: Description,
				Logo:        Logo,
				TokenID:     TokenID,
				MaxSupply:   MaxSupply,
				InitSupply:  InitSupply,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProject(tt.args.projectPB)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProject() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestProject_ToProtobuf(t *testing.T) {
	type fields struct {
		ID          uuid.UUID
		DisplayName string
		Description string
		Logo        string
		TokenID     uuid.UUID
		MaxSupply   float32
		InitSupply  float32
		Links       datatypes.JSON
		Token       token.Token
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.Project
		wantErr bool
	}{
		// Add test cases.
		{
			"happy path",
			fields{ID: ProjectID, DisplayName: DisplayName, Description: Description, Logo: Logo, TokenID: TokenID, MaxSupply: MaxSupply, InitSupply: InitSupply, Links: LinksJSON,
				Token: TokenModel},
			&ProjectPB,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Project{
				ID:          tt.fields.ID,
				DisplayName: tt.fields.DisplayName,
				Description: tt.fields.Description,
				Logo:        tt.fields.Logo,
				TokenID:     tt.fields.TokenID,
				MaxSupply:   tt.fields.MaxSupply,
				InitSupply:  tt.fields.InitSupply,
				Links:       tt.fields.Links,
				Token:       tt.fields.Token,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				DeletedAt:   tt.fields.DeletedAt,
			}
			got, err := p.ToProtobuf()
			if (err != nil) != tt.wantErr {
				t.Errorf("\nProject.ToProtobuf() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("\nProject.ToProtobuf() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
