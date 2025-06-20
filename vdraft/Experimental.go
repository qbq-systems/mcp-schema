package vdraft

type Experimental struct {
	AdditionalProperties Elicitation     `json:"additionalProperties"`
	Description          string          `json:"description"`
	Type                 AnnotationsType `json:"type"`
}
