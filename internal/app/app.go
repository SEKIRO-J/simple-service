package app

import (
	"context"

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
	log.Info("Listing transactions")
	offset := int32(0)
	limit := int32(50)

	if req.PageSize != nil {
		limit = *req.PageSize
	}
	if req.PageToken != nil {
		offset = *req.PageToken
	}

	param := &blockchain.GetAddressInfoParam{Addr: req.Address, Limit: &limit, Offset: &offset}
	txs, err := blockchain.GetAddressInfo(param)
	if err != nil {
		return nil, err
	}

	txPBs := getTransactions(txs)
	nextPageToken := offset + limit

	return &api.ListTransactionsResponse{Transactions: txPBs, NextPageToken: &nextPageToken}, nil
}

func (s *Server) getBalance(ctx context.Context, req *api.GetBalanceRequest) (*api.Balance, error) {

}

func getBalance(rawData *blockchain.TransactionsRawData) float64 {
	return float64(rawData.FinalBalance / blockchain.BaseMultiplier)
}

func getTransactions(rawData *blockchain.TransactionsRawData) []*api.Transaction {
	txsPB := make([]*api.Transaction, len(rawData.Transactions))
	for _, tx := range rawData.Transactions {
		txsPB = append(txsPB, &api.Transaction{Hash: tx.Hash})
	}

	return txsPB
}
