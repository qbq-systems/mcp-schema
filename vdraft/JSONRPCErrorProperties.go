package vdraft

type JSONRPCErrorProperties struct {
	Error   EmptyResult  `json:"error"`
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
}
