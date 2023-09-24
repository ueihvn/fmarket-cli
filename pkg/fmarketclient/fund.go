package fmarketclient

type Fund struct {
	ID               int              `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	ShortName        string           `json:"shortName,omitempty"`
	Type             string           `json:"type,omitempty"`
	NAV              float64          `json:"nav,omitempty"`
	Owner            Owner            `json:"owner,omitempty"`
	FundAssetType    FundAssetType    `json:"dataFundAssetType,omitempty"`
	ManagementFee    float64          `json:"managementFee,omitempty"`
	ProductNavChange ProductNavChange `json:"productNavChange,omitempty"`
	AvgAnnualReturn  float64          `json:"avgAnnualReturn,omitempty"`
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

type ProductNavChange struct {
	NavToPrevious            float64 `json:"navToPrevious,omitempty"`
	NavToLastYear            float64 `json:"navToLastYear,omitempty"`
	NavToEstablish           float64 `json:"navToEstablish,omitempty"`
	NavTo1Months             float64 `json:"navTo1Months,omitempty"`
	NavTo3Months             float64 `json:"navTo3Months,omitempty"`
	NavTo6Months             float64 `json:"navTo6Months,omitempty"`
	NavTo12Months            float64 `json:"navTo12Months,omitempty"`
	NavTo24Months            float64 `json:"navTo24Months,omitempty"`
	NavTo36Months            float64 `json:"navTo36Months,omitempty"`
	AnnualizedReturn36Months float64 `json:"annualizedReturn36Months,omitempty"`
}
