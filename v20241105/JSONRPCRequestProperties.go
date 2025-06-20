package v20241105

type JSONRPCRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc Method          `json:"jsonrpc"`
	Method  Name            `json:"method"`
	Params  HilariousParams `json:"params"`
}
