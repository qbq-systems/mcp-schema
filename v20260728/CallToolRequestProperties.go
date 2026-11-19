package v20260728

type CallToolRequestProperties struct {
	ID      *AdditionalProperties `json:"id,omitempty"`
	Jsonrpc JsonrpcClass          `json:"jsonrpc"`
	Method  JsonrpcClass          `json:"method"`
	Params  AdditionalProperties  `json:"params"`
}
