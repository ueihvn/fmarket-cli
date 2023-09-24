package fmarketclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFmarketClient_FilterFunds(t *testing.T) {
	t.Parallel()
	client, err := NewFmarketClient()
	if err != nil {
		require.NoError(t, err, "new fmarket client %v", err)
	}
	req := NewFilterFundsRequestBuilder().
		SetPage(1).
		SetPageSize(100).
		SetSortOrder(DescSortOrder).
		SetTypes([]string{NewFundType, TradingFundType}).
		SetSortField(AnnualizedReturn36MonthsSortField).
		Build()
	res, err := client.FilterFunds(context.TODO(), &req)
	if err != nil {
		require.NoError(t, err, "filter funds with req: %+v %v", req, err)
	}
	assert.NotEqual(t, res.Total, 0)
	assert.NotEqual(t, len(res.Rows), 0)
}
