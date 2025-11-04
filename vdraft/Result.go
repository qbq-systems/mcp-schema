package vdraft

type Result struct {
	AdditionalProperties AdditionalPropertiesClass    `json:"additionalProperties"`
	Properties           NotificationParamsProperties `json:"properties"`
	Type                 AnnotationsType              `json:"type"`
}
