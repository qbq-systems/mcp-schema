package v20250326

type JSONRPCResponseProperties struct {
	ID      EmptyResult `json:"id"`
	Jsonrpc TypeClass   `json:"jsonrpc"`
	Result  EmptyResult `json:"result"`
}
