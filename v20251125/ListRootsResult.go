package v20251125

type ListRootsResult struct {
	Description string                    `json:"description"`
	Properties  ListRootsResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
