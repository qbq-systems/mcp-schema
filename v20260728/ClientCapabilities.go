package v20260728

type ClientCapabilities struct {
	Description string                       `json:"description"`
	Properties  ClientCapabilitiesProperties `json:"properties"`
	Type        TypeElement                  `json:"type"`
}
