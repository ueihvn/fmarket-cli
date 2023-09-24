package fmarketclient

const (
	NewFundType     = "NEW_FUND"
	TradingFundType = "TRADING_FUND"

	defaultPageSize = 100

	DescSortOrder                     = "DESC"
	AnnualizedReturn36MonthsSortField = "annualizedReturn36Months"

	NavTo1MonthsSortField  = "navTo1Months"
	NavTo3MonthsSortField  = "navTo3Months"
	NavTo6MonthsSortField  = "navTo6Months"
	NavTo12MonthsSortField = "navTo12Months"
	NavTo24MonthsSortField = "navTo24Months"
	NavTo36MonthsSortField = "navTo36Months"
)

// FilterFundsRequestBuilder is a builder for FilterFundsRequest.
type FilterFundsRequestBuilder struct {
	request FilterFundsRequest
}

// NewFilterFundsRequestBuilder creates a new builder instance.
func NewFilterFundsRequestBuilder() *FilterFundsRequestBuilder {
	return &FilterFundsRequestBuilder{}
}

// SetTypes sets the Types field of the request.
func (b *FilterFundsRequestBuilder) SetTypes(types []string) *FilterFundsRequestBuilder {
	b.request.Types = types
	return b
}

// SetIssuerIds sets the IssuerIds field of the request.
func (b *FilterFundsRequestBuilder) SetIssuerIds(issuerIds []int) *FilterFundsRequestBuilder {
	b.request.IssuerIds = issuerIds
	return b
}

// SetPage sets the Page field of the request.
func (b *FilterFundsRequestBuilder) SetPage(page int) *FilterFundsRequestBuilder {
	b.request.Page = page
	return b
}

// SetPageSize sets the PageSize field of the request.
func (b *FilterFundsRequestBuilder) SetPageSize(pageSize int) *FilterFundsRequestBuilder {
	b.request.PageSize = pageSize
	return b
}

// SetSearchField sets the SearchField field of the request.
func (b *FilterFundsRequestBuilder) SetSearchField(searchField string) *FilterFundsRequestBuilder {
	b.request.SearchField = searchField
	return b
}

// SetSortOrder sets the SortOrder field of the request.
func (b *FilterFundsRequestBuilder) SetSortOrder(sortOrder string) *FilterFundsRequestBuilder {
	b.request.SortOrder = sortOrder
	return b
}

// SetSortField sets the SortField field of the request.
func (b *FilterFundsRequestBuilder) SetSortField(sortField string) *FilterFundsRequestBuilder {
	b.request.SortField = sortField
	return b
}

// Build finalizes the construction of the FilterFundsRequest.
func (b *FilterFundsRequestBuilder) Build() FilterFundsRequest {
	if b.request.PageSize == 0 {
		b.SetPageSize(defaultPageSize)
	}
	return b.request
}
