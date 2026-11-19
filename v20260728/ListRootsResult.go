package v20260728

type ListRootsResult struct {
	Description string                    `json:"description"`
	Properties  ListRootsResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        TypeElement               `json:"type"`
}
