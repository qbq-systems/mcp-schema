package vdraft

type JSONRPCNotification struct {
	Description string                        `json:"description"`
	Properties  JSONRPCNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired     `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
