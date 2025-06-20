package v20250326

type ServerCapabilities struct {
	Description string                       `json:"description"`
	Properties  ServerCapabilitiesProperties `json:"properties"`
	Type        AnnotationsType              `json:"type"`
}
