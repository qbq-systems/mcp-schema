package vdraft

type BaseMetadata struct {
	Description string                 `json:"description"`
	Properties  BaseMetadataProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        TypeElement            `json:"type"`
}
