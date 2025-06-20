package v20241105

type Error struct {
	Properties ErrorProperties `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
