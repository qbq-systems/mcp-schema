package v20241105

type ModelPreferences struct {
	Description string                     `json:"description"`
	Properties  ModelPreferencesProperties `json:"properties"`
	Type        AnnotationsType            `json:"type"`
}
