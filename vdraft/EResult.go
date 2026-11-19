package vdraft

type EResult struct {
	Description string                    `json:"description"`
	Properties  CacheableResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        TypeElement               `json:"type"`
}
