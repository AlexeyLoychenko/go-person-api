package model

type WebResponse[T any] struct {
	Data   T             `json:"data,omitempty"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type PageMetadata struct {
	PageId      int  `json:"page_id"`
	HasNextPage bool `json:"has_next_page"`
}
