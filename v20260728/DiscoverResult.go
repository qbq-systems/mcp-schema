package v20260728

type DiscoverResult struct {
	Description string                   `json:"description"`
	Properties  DiscoverResultProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
