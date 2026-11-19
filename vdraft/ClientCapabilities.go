package vdraft

type ClientCapabilities struct {
	Description string                       `json:"description"`
	Properties  ClientCapabilitiesProperties `json:"properties"`
	Type        TypeElement                  `json:"type"`
}
