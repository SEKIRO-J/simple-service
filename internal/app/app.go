package app

import (
	"context"
	"fmt"

	api "github.com/sekiro-j/simpleservice/api/protos/v1"
	"github.com/sekiro-j/simpleservice/internal/db"
	"github.com/sekiro-j/simpleservice/pkg/blockchain"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	dbc *db.Connection
	api.UnimplementedSimpleServiceServer
}

func New(dbc *db.Connection) *Server {
	return &Server{dbc: dbc}
}

func (s *Server) ListTransactions(ctx context.Context, req *api.ListTransactionsRequest) (*api.ListTransactionsResponse, error) {
	log.Info(fmt.Sprintf("Listing transactions for address: %s", req.Address))
	offset := int32(0)
	limit := int32(50)

	if req.PageSize != nil {
		limit = *req.PageSize
	}
	if req.PageToken != nil {
		offset = *req.PageToken
	}

	param := &blockchain.GetAddressInfoParam{Addr: req.Address, Limit: &limit, Offset: &offset}
	rawData, err := blockchain.GetAddressInfo(param)
	if err != nil {
		return nil, err
	}

	txPBs := getTransactions(rawData)
	nextPageToken := offset + limit

	return &api.ListTransactionsResponse{Transactions: txPBs, NextPageToken: nextPageToken}, nil
}

func (s *Server) GetBalance(ctx context.Context, req *api.GetBalanceRequest) (*api.Balance, error) {
	log.Info(fmt.Sprintf("Getting balance for address: %s", req.Address))
	offset := int32(0)
	limit := int32(0)

	param := &blockchain.GetAddressInfoParam{Addr: req.Address, Limit: &limit, Offset: &offset}
	rawData, err := blockchain.GetAddressInfo(param)
	if err != nil {
		return nil, err
	}

	return &api.Balance{FinalBalance: getBalance(rawData), Token: "BTC"}, nil
}

func getBalance(rawData *blockchain.TransactionsRawData) float32 {
	return float32(rawData.FinalBalance / blockchain.BaseMultiplier)
}

func getTransactions(rawData *blockchain.TransactionsRawData) []*api.Transaction {
	txsPB := make([]*api.Transaction, len(rawData.Transactions))
	for _, tx := range rawData.Transactions {
		txsPB = append(txsPB, &api.Transaction{Hash: tx.Hash})
	}

	return txsPB
}
