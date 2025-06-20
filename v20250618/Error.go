package v20250618

type Error struct {
	Properties ErrorProperties `json:"properties"`
	Required   []string        `json:"required"`
	Type       AnnotationsType `json:"type"`
}
