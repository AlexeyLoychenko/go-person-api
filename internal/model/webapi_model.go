package model

type GeneralApiResponse struct {
	ApiKey  string                 `json:"apiname,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Error   error                  `json:"error,omitempty"`
}

type GroupedApiResponse struct {
	Name        string
	Age         int
	Gender      string
	Nationality string
}
