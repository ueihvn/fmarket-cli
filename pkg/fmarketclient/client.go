package fmarketclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type FmarketClient interface {
	FilterFunds(ctx context.Context, req *FilterFundsRequest) (*FilterFundsRespData, error)
}

type fmarketClient struct {
	client *resty.Client
	host   string
}

func NewFmarketClient(opts ...Option) (*fmarketClient, error) {
	cnf := defaultConfig()
	for _, opt := range opts {
		opt(&cnf)
	}
	httpClient := resty.New().SetRetryCount(cnf.retry).SetTimeout(cnf.timeout)
	// set config to client
	return &fmarketClient{
		client: httpClient,
		host:   cnf.host,
	}, nil
}

func (f *fmarketClient) FilterFunds(ctx context.Context, req *FilterFundsRequest) (*FilterFundsRespData, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	url := fmt.Sprintf("%s/%s", f.host, "res/products/filter")
	resp := FmarketResponse[FilterFundsRespData]{}
	res, err := f.client.R().
		SetContext(ctx).
		SetBody(*req).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, err
	}
	defer res.RawResponse.Body.Close()
	data := getData[FilterFundsRespData](resp)
	return &data, nil
}
