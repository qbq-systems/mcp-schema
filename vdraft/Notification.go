package vdraft

type Notification struct {
	Properties NotificationProperties    `json:"properties"`
	Required   []CallToolRequestRequired `json:"required"`
	Type       AnnotationsType           `json:"type"`
}
