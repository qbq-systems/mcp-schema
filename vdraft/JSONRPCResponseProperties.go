package vdraft

type JSONRPCResponseProperties struct {
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Result  EmptyResult  `json:"result"`
}
