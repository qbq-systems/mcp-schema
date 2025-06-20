package vdraft

type JSONRPCResponse struct {
	Description string                    `json:"description"`
	Properties  JSONRPCResponseProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
