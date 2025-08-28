package vdraft

type CompleteRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc JsonrpcClass    `json:"jsonrpc"`
	Method  JsonrpcClass    `json:"method"`
	Params  TentacledParams `json:"params"`
}
