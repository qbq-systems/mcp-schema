package vdraft

type PaginatedRequestProperties struct {
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
	Method  Default              `json:"method"`
	Params  AdditionalProperties `json:"params"`
}
