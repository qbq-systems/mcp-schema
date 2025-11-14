package vdraft

type PurpleTasks struct {
	Description string           `json:"description"`
	Properties  StickyProperties `json:"properties"`
	Type        AnnotationsType  `json:"type"`
}
