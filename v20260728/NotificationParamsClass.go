package v20260728

type NotificationParamsClass struct {
	Description string                       `json:"description"`
	Properties  NotificationParamsProperties `json:"properties"`
	Type        TypeElement                  `json:"type"`
	Required    []string                     `json:"required,omitempty"`
}
