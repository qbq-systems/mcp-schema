package vdraft

type ElicitationCompleteNotificationProperties struct {
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  PurpleParams `json:"params"`
}
