package v20251125

type ElicitationCompleteNotificationProperties struct {
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  PurpleParams `json:"params"`
}
