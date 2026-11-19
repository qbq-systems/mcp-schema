package v20251125

type JSONRPCErrorResponseProperties struct {
	Error   EmptyResult  `json:"error"`
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
}
