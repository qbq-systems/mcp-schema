package v20260728

type FluffyItems struct {
	Description string           `json:"description"`
	Properties  FriskyProperties `json:"properties"`
	Required    []string         `json:"required"`
	Type        TypeElement      `json:"type"`
}
