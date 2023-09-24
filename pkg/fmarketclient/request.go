package fmarketclient

type FilterFundsRequest struct {
	Types       []string `json:"types,omitempty"`
	IssuerIds   []int    `json:"issuerIds,omitempty"`
	Page        int      `json:"page,omitempty"`
	PageSize    int      `json:"pageSize,omitempty"`
	SearchField string   `json:"searchField,omitempty"`
	SortOrder   string   `json:"sortOrder,omitempty"`
	SortField   string   `json:"sortField,omitempty"`
}
