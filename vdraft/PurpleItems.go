package vdraft

type PurpleItems struct {
	Description string           `json:"description"`
	Properties  IndigoProperties `json:"properties"`
	Required    []string         `json:"required"`
	Type        AnnotationsType  `json:"type"`
}
