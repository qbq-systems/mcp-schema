package vdraft

type JSONRPCRequest struct {
	Description string                    `json:"description"`
	Properties  JSONRPCRequestProperties  `json:"properties"`
	Required    []CallToolRequestRequired `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
