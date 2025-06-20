package v20250326

type Error struct {
	Properties ErrorProperties `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
