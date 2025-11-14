package vdraft

type FluffyElicitation struct {
	Description string             `json:"description"`
	Properties  IndecentProperties `json:"properties"`
	Type        AnnotationsType    `json:"type"`
}
