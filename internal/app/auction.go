package app

import (
	"context"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	log "github.com/sirupsen/logrus"
)

func (s *Server) ListAuctions(ctx context.Context, req *api.ListAuctionsRequest) (*api.ListAuctionsResponse, error) {
	log.Info("listing auctions")

	auctions, err := s.dbc.ListAuctions(req.Parent)
	if err != nil {
		return nil, err
	}

	auctionPBs, err := shipyard.BatchConvertAuctions(auctions)
	if err != nil {
		return nil, err
	}

	return &api.ListAuctionsResponse{Auctions: auctionPBs}, nil
}

func (s *Server) CreateAuction(ctx context.Context, req *api.CreateAuctionRequest) (*api.Auction, error) {
	log.WithFields(log.Fields{
		"request": req,
	}).Info("creating auction")

	auction, err := s.dbc.CreateAuction(req.Parent, req.GetAuction())
	if err != nil {
		return nil, err
	}

	auctionPB, err := auction.ToProtobuf()
	if err != nil {
		return nil, err
	}

	return auctionPB, nil
}
