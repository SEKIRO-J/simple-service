package shipyard

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	TokenDisplayName string = "Green Metaverse Token"
	TokenSymbol      string = "GMT"
	TokenLogo        string = "http://www.cdn.com/token"
	TokenAddr        string = "0x12345678"

	FundTokenDisplayName string = "Fund Token"
	FundTokenSymbol      string = "FT"
	FundTokenLogo        string = "http://www.cdn.com/fundtoken"
	FundTokenAddr        string = "0x12345677"
)

var (
	TokenID    = uuid.New()
	TokenIDstr = TokenID.String()

	TokenPB = api.Token{
		Name:        tokens + TokenIDstr,
		DisplayName: TokenDisplayName,
		Symbol:      TokenSymbol,
		Logo:        TokenLogo,
		Addr:        TokenAddr,
	}

	TokenModel = token.Token{
		ID:          TokenID,
		DisplayName: TokenDisplayName,
		Symbol:      TokenSymbol,
		Logo:        TokenLogo,
		Addr:        TokenAddr,
	}

	FundTokenID    = uuid.New()
	FundTokenIDstr = FundTokenID.String()

	FundTokenPB = api.Token{
		Name:        tokens + FundTokenIDstr,
		DisplayName: FundTokenDisplayName,
		Symbol:      FundTokenSymbol,
		Logo:        FundTokenLogo,
		Addr:        FundTokenAddr,
	}

	FundTokenModel = token.Token{
		ID:          FundTokenID,
		DisplayName: FundTokenDisplayName,
		Symbol:      FundTokenSymbol,
		Logo:        FundTokenLogo,
		Addr:        FundTokenAddr,
	}
)

var (
	AuctionID    = uuid.New()
	AuctionIDstr = AuctionID.String()

	Stages map[string]interface{} = map[string]interface{}{
		"kyc": "vist https://www.xxxx.com for kyc details",
	}
	StagesPB, _   = structpb.NewStruct(Stages)
	StagesJSON, _ = protojson.Marshal(StagesPB)
)

const (
	PoolID          string  = "1"
	Crowdsale       string  = "crowdsale"
	Dutch           string  = "dutch-auction"
	OfferAmount     float32 = 12345678.12345678
	RemainingAmount float32 = 12345678.12345678
	TargetAmount    float32 = 345678.12345678
)

func TestNewAuction(t *testing.T) {
	type args struct {
		auctionPB *api.Auction
	}
	tests := []struct {
		name    string
		args    args
		want    *Auction
		wantErr bool
	}{
		// Add test cases.
		{
			"happy path create with token id",
			args{
				&api.Auction{
					Name:        "",
					ProjectId:   ProjectIDstr,
					PoolId:      PoolID,
					SaleType:    Crowdsale,
					Offer:       OfferAmount,
					Remaining:   RemainingAmount,
					FundTokenId: FundTokenIDstr,
					Target:      TargetAmount,
				},
			},
			&Auction{
				ProjectID:   ProjectID,
				PoolID:      PoolID,
				SaleType:    Crowdsale,
				Offer:       OfferAmount,
				Remaining:   RemainingAmount,
				FundTokenID: FundTokenID,
				Target:      TargetAmount,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuction(tt.args.auctionPB)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuction() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestAuction_ToProtobuf(t *testing.T) {
	type fields struct {
		ID          uuid.UUID
		ProjectID   uuid.UUID
		PoolID      string
		SaleType    string
		Offer       float32
		Remaining   float32
		FundTokenID uuid.UUID
		Target      float32
		Stages      datatypes.JSON
		Project     Project
		FundToken   token.Token
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.Auction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"happy path",
			fields{
				ID:          AuctionID,
				ProjectID:   ProjectID,
				PoolID:      PoolID,
				SaleType:    Crowdsale,
				Offer:       OfferAmount,
				Remaining:   RemainingAmount,
				FundTokenID: FundTokenID,
				Target:      TargetAmount,
				Stages:      StagesJSON,
				Project:     ProjectModel,
				FundToken:   FundTokenModel,
			},
			&api.Auction{
				Name:        projects + ProjectIDstr + delimiter + auctions + AuctionIDstr,
				ProjectId:   ProjectIDstr,
				PoolId:      PoolID,
				SaleType:    Crowdsale,
				Offer:       OfferAmount,
				Remaining:   RemainingAmount,
				FundTokenId: FundTokenIDstr,
				Target:      TargetAmount,
				FundToken:   &FundTokenPB,
				Project:     &ProjectPB,
				Stages:      StagesPB,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Auction{
				ID:          tt.fields.ID,
				ProjectID:   tt.fields.ProjectID,
				PoolID:      tt.fields.PoolID,
				SaleType:    tt.fields.SaleType,
				Offer:       tt.fields.Offer,
				Remaining:   tt.fields.Remaining,
				FundTokenID: tt.fields.FundTokenID,
				Target:      tt.fields.Target,
				Stages:      tt.fields.Stages,
				Project:     tt.fields.Project,
				FundToken:   tt.fields.FundToken,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				DeletedAt:   tt.fields.DeletedAt,
			}
			got, err := a.ToProtobuf()
			if (err != nil) != tt.wantErr {
				t.Errorf("Auction.ToProtobuf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("Auction.ToProtobuf() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
