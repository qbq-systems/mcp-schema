package v20260728

type SubscriptionFilter struct {
	Description string                       `json:"description"`
	Properties  SubscriptionFilterProperties `json:"properties"`
	Type        TypeElement                  `json:"type"`
}
