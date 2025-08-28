package vdraft

type JSONRPCNotificationProperties struct {
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  Default      `json:"method"`
	Params  Result       `json:"params"`
}
