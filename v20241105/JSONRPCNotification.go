package v20241105

type JSONRPCNotification struct {
	Description string                        `json:"description"`
	Properties  JSONRPCNotificationProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
