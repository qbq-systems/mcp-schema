package vdraft

type ProgressNotificationProperties struct {
	Jsonrpc JsonrpcClass  `json:"jsonrpc"`
	Method  JsonrpcClass  `json:"method"`
	Params  MagentaParams `json:"params"`
}
