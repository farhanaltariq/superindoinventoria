package models

type Pagination struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalRows  int    `json:"totalRows"`
	TotalPages int    `json:"totalPages"`
	Filter     string `json:"filter"`
	Sort       string `json:"sort"`
	Dir        string `json:"dir"`
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	TotalRows  int         `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Page       int         `json:"page"`
}
