package vdraft

type JSONRPCErrorResponse struct {
	Description string                         `json:"description"`
	Properties  JSONRPCErrorResponseProperties `json:"properties"`
	Required    []string                       `json:"required"`
	Type        TypeElement                    `json:"type"`
}
