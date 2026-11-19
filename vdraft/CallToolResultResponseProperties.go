package vdraft

type CallToolResultResponseProperties struct {
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
	Result  ClientRequest        `json:"result"`
}
