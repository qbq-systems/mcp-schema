package v20260728

type HeaderMismatchErrorProperties struct {
	Error   PurpleError          `json:"error"`
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
}
