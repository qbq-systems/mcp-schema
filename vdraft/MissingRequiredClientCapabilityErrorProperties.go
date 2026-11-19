package vdraft

type MissingRequiredClientCapabilityErrorProperties struct {
	Error   FluffyError          `json:"error"`
	ID      AdditionalProperties `json:"id"`
	Jsonrpc JsonrpcClass         `json:"jsonrpc"`
}
