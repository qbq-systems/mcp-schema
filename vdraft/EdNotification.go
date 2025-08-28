package vdraft

type EdNotification struct {
	Description string                            `json:"description"`
	Properties  InitializedNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired         `json:"required"`
	Type        AnnotationsType                   `json:"type"`
}
