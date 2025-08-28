package vdraft

type JSONRPCRequestProperties struct {
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  Default      `json:"method"`
	Params  Result       `json:"params"`
}
