package fmarketclient

type FmarketResponse[K any] struct {
	Code    int    `json:"code,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    K      `json:"data,omitempty"`
}

func getData[K any](resp FmarketResponse[K]) K {
	return resp.Data
}
