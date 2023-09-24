package fmarketclient

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestFmarketClient_FilterFunds(t *testing.T) {
	client, err := NewFmarketClient()
	if err != nil {
		t.Errorf("want nil but got: %v", err)
	}
	req := FilterFundsRequest{
		Types: []string{
			"NEW_FUND",
			"TRADING_FUND",
		},
		Page:        1,
		PageSize:    100,
		SearchField: "",
		SortOrder:   "DESC",
	}
	res, err := client.FilterFunds(context.TODO(), &req)
	if err != nil {
		t.Errorf("want nil but got: %v", err)
	}
	log.Info(res)
}
