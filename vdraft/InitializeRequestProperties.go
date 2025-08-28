package vdraft

type InitializeRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc JsonrpcClass    `json:"jsonrpc"`
	Method  JsonrpcClass    `json:"method"`
	Params  HilariousParams `json:"params"`
}
