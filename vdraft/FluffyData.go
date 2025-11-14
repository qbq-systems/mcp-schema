package vdraft

type FluffyData struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Properties           Properties4               `json:"properties"`
	Required             []string                  `json:"required"`
	Type                 AnnotationsType           `json:"type"`
}
