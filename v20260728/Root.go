package v20260728

type Root struct {
	Description string          `json:"description"`
	Properties  RootProperties  `json:"properties"`
	Required    []FormatElement `json:"required"`
	Type        TypeElement     `json:"type"`
}
