package response

// ProxyResponse Represents the JSON response body
type ProxyResponse struct {
	ID      int               `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int64             `json:"length"`
}
