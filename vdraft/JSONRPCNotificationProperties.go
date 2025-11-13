package vdraft

type JSONRPCNotificationProperties struct {
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
	Method  Default      `json:"method"`
	Params  FluffyParams `json:"params"`
	ID      *EmptyResult `json:"id,omitempty"`
}
