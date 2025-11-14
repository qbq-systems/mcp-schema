package vdraft

type Experimental struct {
	AdditionalProperties Metadata        `json:"additionalProperties"`
	Description          string          `json:"description"`
	Type                 AnnotationsType `json:"type"`
}
