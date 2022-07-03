package app

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"github.com/golang/protobuf/jsonpb"
	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListEvents(ctx context.Context, req *api.ListEventsRequest) (*api.ListEventsResponse, error) {
	log.Info("listing events")

	events, err := s.dbc.ListEvents()
	if err != nil {
		return nil, err
	}

	auctionPBs, err := shipyard.BatchConvertEvents(events)
	if err != nil {
		return nil, err
	}

	return &api.ListEventsResponse{Events: auctionPBs}, nil
}

func (s *Server) BatchCreateEvents(ctx context.Context, req *api.BatchCreateEventsRequest) (*api.BatchCreateEventsResponse, error) {
	log.WithFields(log.Fields{
		"request": req,
	}).Info("creating event")

	marshaler := jsonpb.Marshaler{EmitDefaults: true}
	message, _ := marshaler.MarshalToString(req.Payload)

	if req.Hmac == nil || !checkAuth(string(message), req.Hmac.Nonce, s.flowScannerKey, req.Hmac.Hash) {
		errMSG := "permission denied"
		log.Error(errMSG)
		return nil, status.Error(codes.PermissionDenied, errMSG)
	}

	events, err := s.dbc.CreateEvents(req.Payload)
	if err != nil {
		return nil, err
	}

	eventPBs, err := shipyard.BatchConvertEvents(events)
	if err != nil {
		return nil, err
	}

	return &api.BatchCreateEventsResponse{Events: eventPBs}, nil
}

func checkAuth(message string, nonce string, key string, expectedHMAC string) bool {
	h := sha256.New()
	h.Write([]byte(nonce + message))
	hashDigest := hex.EncodeToString(h.Sum(nil))

	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(message + hashDigest))
	hmacDigest := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return hmacDigest == expectedHMAC
}
