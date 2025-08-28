package vdraft

type ListRootsRequestProperties struct {
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  Result       `json:"params"`
}
