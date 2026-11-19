package v20260728

type Completion struct {
	Properties CompletionProperties `json:"properties"`
	Required   []string             `json:"required"`
	Type       TypeElement          `json:"type"`
}
