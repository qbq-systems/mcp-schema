package vdraft

type ReadResourceRequestProperties struct {
	ID      *EmptyResult `json:"id,omitempty"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  FriskyParams `json:"params"`
}
