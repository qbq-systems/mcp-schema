package v20250326

type ModelHint struct {
	Description string              `json:"description"`
	Properties  ModelHintProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
