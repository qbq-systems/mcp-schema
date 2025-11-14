package vdraft

type FluffyData struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Properties           AmbitiousProperties       `json:"properties"`
	Required             []string                  `json:"required"`
	Type                 AnnotationsType           `json:"type"`
}
