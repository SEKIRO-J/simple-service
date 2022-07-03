package db

import (
	"fmt"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Connection) CreateEvents(eventPB *api.EventsPayload) ([]*shipyard.Event, error) {
	events, err := shipyard.NewEvents(eventPB)
	if err != nil {
		return nil, err
	}

	if err := c.orm.Create(events).Error; err != nil {
		errMsg := fmt.Sprintf("failed to save events: %v", err)
		log.WithFields(log.Fields{
			"events": events,
		}).Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return events, nil
}

func (c *Connection) ListEvents() ([]*shipyard.Event, error) {
	var events []*shipyard.Event

	if err := c.orm.Find(&events).Error; err != nil {
		errMsg := fmt.Sprintf("failed to list events: %v", err)
		log.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return events, nil
}
