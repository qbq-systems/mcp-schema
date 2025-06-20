package v20241105

type JSONRPCRequest struct {
	Description string                   `json:"description"`
	Properties  JSONRPCRequestProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
