package vdraft

type PurpleElicitation struct {
	Description string           `json:"description"`
	Properties  FluffyProperties `json:"properties"`
	Type        AnnotationsType  `json:"type"`
}
