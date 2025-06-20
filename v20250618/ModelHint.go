package v20250618

type ModelHint struct {
	Description string              `json:"description"`
	Properties  ModelHintProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
