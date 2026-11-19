package vdraft

type Notification struct {
	Properties NotificationProperties    `json:"properties"`
	Required   []CallToolRequestRequired `json:"required"`
	Type       TypeElement               `json:"type"`
}
