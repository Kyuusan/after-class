package dto


type ResponseWrapper [T any] struct {
	Data *T `json:"data"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Pagination *PaginationResponse `json:"paginantion"`
} 