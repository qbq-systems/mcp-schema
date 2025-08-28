package vdraft

type CreateMessageRequestProperties struct {
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  StickyParams `json:"params"`
}
