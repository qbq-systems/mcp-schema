package vdraft

type CallToolRequestProperties struct {
	ID      *EmptyResult `json:"id,omitempty"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  EmptyResult  `json:"params"`
}
