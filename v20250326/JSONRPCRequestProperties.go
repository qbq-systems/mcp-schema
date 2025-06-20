package v20250326

type JSONRPCRequestProperties struct {
	ID      EmptyResult     `json:"id"`
	Jsonrpc TypeClass       `json:"jsonrpc"`
	Method  Name            `json:"method"`
	Params  HilariousParams `json:"params"`
}
