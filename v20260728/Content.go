package v20260728

type Content struct {
	Description string                 `json:"description"`
	Properties  AudioContentProperties `json:"properties"`
	Required    []string               `json:"required"`
	Type        TypeElement            `json:"type"`
}
