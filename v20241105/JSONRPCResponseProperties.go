package v20241105

type JSONRPCResponseProperties struct {
	ID      EmptyResult `json:"id"`
	Jsonrpc Method      `json:"jsonrpc"`
	Result  EmptyResult `json:"result"`
}
