package v20260728

type ProgressNotificationParams struct {
	Description string                               `json:"description"`
	Properties  ProgressNotificationParamsProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        TypeElement                          `json:"type"`
}
