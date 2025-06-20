package v20241105

type JSONRPCErrorProperties struct {
	Error   Error       `json:"error"`
	ID      EmptyResult `json:"id"`
	Jsonrpc Method      `json:"jsonrpc"`
}
