package v20250618

type JSONRPCNotificationProperties struct {
	Jsonrpc MethodClass `json:"jsonrpc"`
	Method  Default     `json:"method"`
	Params  Result      `json:"params"`
}
