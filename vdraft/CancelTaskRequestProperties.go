package vdraft

type CancelTaskRequestProperties struct {
	ID      EmptyResult         `json:"id"`
	Jsonrpc JsonrpcClass        `json:"jsonrpc"`
	Method  JsonrpcClass        `json:"method"`
	Params  RelatedTaskMetadata `json:"params"`
}
