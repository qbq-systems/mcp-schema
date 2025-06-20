package v20250326

type JSONRPCMessage struct {
	AnyOf       []JsonrpcBatchRequest `json:"anyOf"`
	Description string                `json:"description"`
}
