package v20260728

type Notification struct {
	Properties NotificationProperties    `json:"properties"`
	Required   []CallToolRequestRequired `json:"required"`
	Type       TypeElement               `json:"type"`
}
