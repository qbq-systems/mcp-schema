package v20260728

type Jsonrpc struct {
	Description string                        `json:"description"`
	Properties  JSONRPCNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired     `json:"required"`
	Type        TypeElement                   `json:"type"`
}
