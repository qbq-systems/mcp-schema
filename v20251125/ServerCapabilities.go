package v20251125

type ServerCapabilities struct {
	Description string                       `json:"description"`
	Properties  ServerCapabilitiesProperties `json:"properties"`
	Type        AnnotationsType              `json:"type"`
}
