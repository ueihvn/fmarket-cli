package fmarketclient

type Fund struct {
	ID            int           `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	ShortName     string        `json:"shortName,omitempty"`
	Type          string        `json:"type,omitempty"`
	NAV           float64       `json:"nav,omitempty"`
	Owner         Owner         `json:"owner,omitempty"`
	FundAssetType FundAssetType `json:"dataFundAssetType,omitempty"`
}

type FundAssetType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type Owner struct {
	ID        int    `json:"id,omitempty"`
	ShortName string `json:"shortName,omitempty"`
}

type FilterFundsRespData struct {
	Rows  []Fund   `json:"rows,omitempty"`
	Total int      `json:"total,omitempty"`
	Types []string `json:"types,omitempty"`
}
