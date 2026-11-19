package vdraft

type HeaderMismatchErrorProperties struct {
	Error   PurpleError          `json:"error"`
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
}
