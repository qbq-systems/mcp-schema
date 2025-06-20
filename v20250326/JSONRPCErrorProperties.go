package v20250326

type JSONRPCErrorProperties struct {
	Error   Error       `json:"error"`
	ID      EmptyResult `json:"id"`
	Jsonrpc TypeClass   `json:"jsonrpc"`
}
