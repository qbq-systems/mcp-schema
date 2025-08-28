package vdraft

type PaginatedRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc JsonrpcClass    `json:"jsonrpc"`
	Method  Default         `json:"method"`
	Params  AmbitiousParams `json:"params"`
}
