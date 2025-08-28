package vdraft

type CancelledNotificationProperties struct {
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  JsonrpcClass `json:"method"`
	Params  FluffyParams `json:"params"`
}
