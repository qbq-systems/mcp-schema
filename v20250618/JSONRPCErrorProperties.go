package v20250618

type JSONRPCErrorProperties struct {
	Error   Error       `json:"error"`
	ID      EmptyResult `json:"id"`
	Jsonrpc MethodClass `json:"jsonrpc"`
}
