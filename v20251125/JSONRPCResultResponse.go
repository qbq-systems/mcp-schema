package v20251125

type JSONRPCResultResponse struct {
	Description string                          `json:"description"`
	Properties  JSONRPCResultResponseProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
