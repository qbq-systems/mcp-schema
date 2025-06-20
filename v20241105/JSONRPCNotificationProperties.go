package v20241105

type JSONRPCNotificationProperties struct {
	Jsonrpc Method `json:"jsonrpc"`
	Method  Name   `json:"method"`
	Params  Result `json:"params"`
}
