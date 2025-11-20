package vdraft

type FluffyTasks struct {
	Description string           `json:"description"`
	Properties  FriskyProperties `json:"properties"`
	Type        AnnotationsType  `json:"type"`
}
