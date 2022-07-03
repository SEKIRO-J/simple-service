package shipyard

import (
	"encoding/json"
	"fmt"
	"time"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type Event struct {
	ID               uuid.UUID      `gorm:"primaryKey; default:uuid_generate_v4()" json:"-"`
	BlockID          string         `json:"blockId"`
	BlockHeight      float32        `json:"blockHeight"`
	BlockTimestamp   time.Time      `json:"blockTimestamp"`
	Type             string         `gorm:"index" json:"type"`
	TransactionID    string         `json:"transactionId"`
	TransactionIndex int            `json:"transactionIndex"`
	EventIndex       int            `json:"eventIndex"`
	Data             datatypes.JSON `json:"data"`
	PoolID           string         `gorm:"index" json:"-"`
	Address          string         `gorm:"index" json:"-"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}

func (Event) TableName() string {
	return "shipyard_events"
}

func (e *Event) ToProtobuf() (*api.Event, error) {
	eventPB := new(api.Event)

	eventJSON, err := json.Marshal(e)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal event: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}
	err = protojson.Unmarshal(eventJSON, eventPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal event: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return eventPB, nil
}

func BatchConvertEvents(events []*Event) ([]*api.Event, error) {
	eventPBs := make([]*api.Event, len(events))
	for i, event := range events {
		eventPB, err := event.ToProtobuf()
		if err != nil {
			return eventPBs, err
		}
		eventPBs[i] = eventPB
	}

	return eventPBs, nil
}

func NewEvent(eventPB *api.Event) (*Event, error) {
	eventJSON, err := protojson.Marshal(eventPB)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal event protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	event := new(Event)
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal event protobuf: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	eventData := new(map[string]interface{})
	err = json.Unmarshal(event.Data, eventData)
	if err != nil {
		errMsg := fmt.Sprintf("failed to unmarshal event data: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	if poolID, exist := (*eventData)["poolId"]; exist {
		event.PoolID = fmt.Sprintf("%v", poolID)
	}
	if address, exist := (*eventData)["address"]; exist {
		event.Address = fmt.Sprintf("%v", address)
	}

	return event, nil
}

func NewEvents(eventsPayloadPB *api.EventsPayload) ([]*Event, error) {
	eventsPB := eventsPayloadPB.Data.GetEvents()

	count := len(eventsPB)
	events := make([]*Event, count)
	var err error
	for i, eventPB := range eventsPB {
		events[i], err = NewEvent(eventPB)
		if err != nil {
			return nil, nil
		}
	}

	return events, nil
}
