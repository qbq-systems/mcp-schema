package vdraft

type JSONRPCErrorResponseProperties struct {
	Error   AdditionalProperties `json:"error"`
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
}
