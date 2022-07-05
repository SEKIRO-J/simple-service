package blockchain

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

const (
	BaseMultiplier        = float32(100000000)
	SingleAddressEndpoint = "https://blockchain.info/rawaddr/%s"
)

type GetAddressInfoParam struct {
	Addr   string
	Limit  *int32
	Offset *int32
}

type TransactionsRawData struct {
	Transactions []Transaction `json:"txs"`
	FinalBalance int32         `json:"final_balance"`
}

type Transaction struct {
	Hash string `json:"hash"`
}

func GetAddressInfo(param *GetAddressInfoParam) (*TransactionsRawData, error) {
	url := fmt.Sprintf(SingleAddressEndpoint, param.Addr)

	urlParam := fmt.Sprintf("?limit=%v&offset=%v", *param.Limit, *param.Offset)
	log.Info(fmt.Sprintf("Calling blockchain.info endpoint: %s", url+urlParam))

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := c.Get(url + urlParam)
	if err != nil {
		log.Error(err)
	}
	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		log.Error(fmt.Printf("status code: %v", resp.Status))
		return nil, status.Error(codes.ResourceExhausted, "blockchain.info is currently not available")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "unable to read response body")
	}
	// log.Info(fmt.Printf("Body : %v", string(body)))

	rawData := &TransactionsRawData{}
	err = json.Unmarshal([]byte(body), rawData)

	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "unable to unmarshal response body")
	}

	return rawData, nil
}
