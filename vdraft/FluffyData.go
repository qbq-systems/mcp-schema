package vdraft

type FluffyData struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Properties           HilariousProperties       `json:"properties"`
	Required             []string                  `json:"required"`
	Type                 AnnotationsType           `json:"type"`
}
