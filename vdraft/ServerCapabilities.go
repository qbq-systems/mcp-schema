package vdraft

type ServerCapabilities struct {
	Description string                       `json:"description"`
	Properties  ServerCapabilitiesProperties `json:"properties"`
	Type        TypeElement                  `json:"type"`
}
