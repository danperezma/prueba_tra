package zincsearch

type CreateDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

type SearchDocumentsRequest struct {
	SearchType string                      `json:"search_type"`
	MaxResults int                         `json:"max_results"`
	Query      SearchDocumentsRequestQuery `json:"query"`
	// SortField string[]					   `json:"sort_fields"`
}

type SearchDocumentsRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}