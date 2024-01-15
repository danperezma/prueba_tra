package zincsearch

// CreateDocumentsRequest is the request for the CreateDocuments function
type CreateDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

// SearchDocumentsRequest is the request for the SearchDocuments function
type SearchDocumentsRequest struct {
	SearchType string                      `json:"search_type"`
	MaxResults int                         `json:"max_results"`
	Query      SearchDocumentsRequestQuery `json:"query"`
}

// SearchDocumentsRequestQuery is the query for the SearchDocuments function
type SearchDocumentsRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}