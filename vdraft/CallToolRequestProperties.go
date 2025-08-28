package vdraft

type CallToolRequestProperties struct {
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  PurpleParams `json:"params"`
}
