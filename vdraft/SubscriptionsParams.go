package vdraft

type SubscriptionsParams struct {
	Description string                                                `json:"description"`
	Properties  SubscriptionsAcknowledgedNotificationParamsProperties `json:"properties"`
	Required    []string                                              `json:"required"`
	Type        TypeElement                                           `json:"type"`
}
