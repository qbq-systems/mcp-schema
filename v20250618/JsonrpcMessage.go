package v20250618

type JsonrpcMessage struct {
	AnyOf       []EmptyResult `json:"anyOf"`
	Description string        `json:"description"`
}
