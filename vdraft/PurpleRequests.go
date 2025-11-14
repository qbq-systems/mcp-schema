package vdraft

type PurpleRequests struct {
	Description string           `json:"description"`
	Properties  IndigoProperties `json:"properties"`
	Type        AnnotationsType  `json:"type"`
}
