package vdraft

type UnsupportedProtocolVersionErrorProperties struct {
	Error   TentacledError       `json:"error"`
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
}
