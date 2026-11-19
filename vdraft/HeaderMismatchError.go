package vdraft

type HeaderMismatchError struct {
	Description string                        `json:"description"`
	Properties  HeaderMismatchErrorProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        TypeElement                   `json:"type"`
}
