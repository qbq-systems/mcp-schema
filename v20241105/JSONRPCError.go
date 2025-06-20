package v20241105

type JSONRPCError struct {
	Description string                 `json:"description"`
	Properties  JSONRPCErrorProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        AnnotationsType        `json:"type"`
}
