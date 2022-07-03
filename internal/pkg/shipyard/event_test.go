package shipyard

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	TimestampEpoch = 1652678319
	BlockID        = "fe0fa8fed08b14f457d8be07c7fdea6a45be83c74dbf3e6276fc316b40eb1b9f"
	BlockHeight    = 29844894
	EventType      = "A.c1e4f4f4c4257510.TopShotMarketV3.MomentWithdrawn"
	TxnID          = "b0644c29667545fce16aa160c25364a70c39794919ec3119ae176a7e6eff2804"
	TxnIndex       = 0
	EventIndex     = 0
	poolID         = "30"
	Address        = "0x08c2ce90cf8a0b68"
)

var (
	BlockTimestamp, _ = time.Parse(time.RFC3339, "2022-05-15T23:22:11.892Z")
	BlockTimestampPB  = timestamppb.New(BlockTimestamp)

	EventData map[string]interface{} = map[string]interface{}{
		"poolId":  poolID,
		"address": Address,
	}

	EventDataPB, _   = structpb.NewStruct(EventData)
	EventDataJSON, _ = protojson.Marshal(EventDataPB)

	EventsPB []*api.Event = []*api.Event{
		{
			BlockId:          BlockID,
			BlockHeight:      BlockHeight,
			BlockTimestamp:   BlockTimestampPB,
			Type:             EventType,
			TransactionId:    TxnID,
			TransactionIndex: TxnIndex,
			EventIndex:       EventIndex,
			Data:             EventDataPB,
		},
	}

	EventsPayloadDataPB *api.EventData = &api.EventData{
		TransactionId: TxnID,
		BlockHeight:   BlockHeight,
		Events:        EventsPB,
	}

	EventsPayloadPB *api.EventsPayload = &api.EventsPayload{
		Timestamp: TimestampEpoch,
		Data:      EventsPayloadDataPB,
	}
)

func TestNewEvents(t *testing.T) {
	type args struct {
		eventPB *api.EventsPayload
	}
	tests := []struct {
		name    string
		args    args
		want    []*Event
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"happy path",
			args{
				&api.EventsPayload{
					Timestamp: TimestampEpoch,
					Data:      EventsPayloadDataPB,
				},
			},
			[]*Event{
				{
					BlockID:          BlockID,
					BlockHeight:      BlockHeight,
					BlockTimestamp:   BlockTimestamp,
					Type:             EventType,
					TransactionID:    TxnID,
					TransactionIndex: TxnIndex,
					EventIndex:       EventIndex,
					Data:             EventDataJSON,
					PoolID:           poolID,
					Address:          Address,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEvents(tt.args.eventPB)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvents() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestEvent_ToProtobuf(t *testing.T) {
	eventID := uuid.New()
	type fields struct {
		ID               uuid.UUID
		BlockID          string
		BlockHeight      float32
		BlockTimestamp   time.Time
		Type             string
		TransactionID    string
		TransactionIndex int
		EventIndex       int
		Data             datatypes.JSON
		PoolID           string
		Address          string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		DeletedAt        gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.Event
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"happy path",
			fields{
				ID:               eventID,
				BlockID:          BlockID,
				BlockHeight:      BlockHeight,
				BlockTimestamp:   BlockTimestamp,
				Type:             EventType,
				TransactionID:    TxnID,
				TransactionIndex: TxnIndex,
				EventIndex:       EventIndex,
				Data:             EventDataJSON,
				PoolID:           poolID,
				Address:          Address,
			},
			&api.Event{
				BlockId:          BlockID,
				BlockHeight:      BlockHeight,
				BlockTimestamp:   BlockTimestampPB,
				Type:             EventType,
				TransactionId:    TxnID,
				TransactionIndex: TxnIndex,
				EventIndex:       EventIndex,
				Data:             EventDataPB,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				ID:               tt.fields.ID,
				BlockID:          tt.fields.BlockID,
				BlockHeight:      tt.fields.BlockHeight,
				BlockTimestamp:   tt.fields.BlockTimestamp,
				Type:             tt.fields.Type,
				TransactionID:    tt.fields.TransactionID,
				TransactionIndex: tt.fields.TransactionIndex,
				EventIndex:       tt.fields.EventIndex,
				Data:             tt.fields.Data,
				PoolID:           tt.fields.PoolID,
				Address:          tt.fields.Address,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				DeletedAt:        tt.fields.DeletedAt,
			}
			got, err := e.ToProtobuf()
			if (err != nil) != tt.wantErr {
				t.Errorf("Event.ToProtobuf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("Event.ToProtobuf() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
