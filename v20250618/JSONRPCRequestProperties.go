package v20250618

type JSONRPCRequestProperties struct {
	ID      EmptyResult `json:"id"`
	Jsonrpc MethodClass `json:"jsonrpc"`
	Method  Default     `json:"method"`
	Params  Result      `json:"params"`
}
