package v20250326

type ClientCapabilities struct {
	Description string                       `json:"description"`
	Properties  ClientCapabilitiesProperties `json:"properties"`
	Type        AnnotationsType              `json:"type"`
}
