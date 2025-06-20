package v20241105

type EmbeddedResourceProperties struct {
	Annotations Annotations        `json:"annotations"`
	Resource    ClientNotification `json:"resource"`
	Type        Method             `json:"type"`
}
