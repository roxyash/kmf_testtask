package model

type ProxyModel struct {
	ID      int               `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int64             `json:"length"`
}
