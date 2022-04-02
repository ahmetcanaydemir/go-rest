package model

type InMemoryRequest struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type InMemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
