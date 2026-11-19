package v20260728

type JSONRPCNotificationProperties struct {
	Jsonrpc JsonrpcClass          `json:"jsonrpc"`
	Method  Default               `json:"method"`
	Params  PurpleParams          `json:"params"`
	ID      *AdditionalProperties `json:"id,omitempty"`
}
