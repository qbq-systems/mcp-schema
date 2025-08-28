package vdraft

type LoggingMessageNotificationProperties struct {
	Jsonrpc JsonrpcClass  `json:"jsonrpc"`
	Method  JsonrpcClass  `json:"method"`
	Params  CunningParams `json:"params"`
}
