package token

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
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

	TokenModel = Token{
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

	FundTokenModel = Token{
		ID:          FundTokenID,
		DisplayName: FundTokenDisplayName,
		Symbol:      FundTokenSymbol,
		Logo:        FundTokenLogo,
		Addr:        FundTokenAddr,
	}
)

func TestToken_ToProtobuf(t *testing.T) {
	type fields struct {
		ID          uuid.UUID
		DisplayName string
		Symbol      string
		Logo        string
		Addr        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Token{
				ID:          tt.fields.ID,
				DisplayName: tt.fields.DisplayName,
				Symbol:      tt.fields.Symbol,
				Logo:        tt.fields.Logo,
				Addr:        tt.fields.Addr,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				DeletedAt:   tt.fields.DeletedAt,
			}
			got, err := tr.ToProtobuf()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token.ToProtobuf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Token.ToProtobuf() = %v, want %v", got, tt.want)
			}
		})
	}
}
