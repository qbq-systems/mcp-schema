package vdraft

type InputSchema struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Description          string                    `json:"description"`
	Properties           IndigoProperties          `json:"properties"`
	Required             []string                  `json:"required"`
	Type                 TypeElement               `json:"type"`
}
