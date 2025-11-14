package vdraft

type PurpleItems struct {
	Description string             `json:"description"`
	Properties  IndecentProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        AnnotationsType    `json:"type"`
}
