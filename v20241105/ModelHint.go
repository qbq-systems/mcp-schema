package v20241105

type ModelHint struct {
	Description string              `json:"description"`
	Properties  ModelHintProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
