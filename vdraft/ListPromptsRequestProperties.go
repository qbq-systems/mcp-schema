package vdraft

type ListPromptsRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc JsonrpcClass    `json:"jsonrpc"`
	Method  JsonrpcClass    `json:"method"`
	Params  AmbitiousParams `json:"params"`
}
