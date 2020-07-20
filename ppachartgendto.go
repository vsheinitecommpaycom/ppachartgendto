package dto

type DataItemDto struct {
	Name       string  `json:"name"`
	Count      uint32  `json:"count"`
	Percentile float64 `json:"percentile"`
}
