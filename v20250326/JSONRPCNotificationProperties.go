package v20250326

type JSONRPCNotificationProperties struct {
	Jsonrpc TypeClass `json:"jsonrpc"`
	Method  Name      `json:"method"`
	Params  Result    `json:"params"`
}
