package vdraft

type UnsupportedProtocolVersionError struct {
	Description string                                    `json:"description"`
	Properties  UnsupportedProtocolVersionErrorProperties `json:"properties"`
	Required    []string                                  `json:"required"`
	Type        TypeElement                               `json:"type"`
}
