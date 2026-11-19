package v20260728

type Error struct {
	Properties ErrorProperties `json:"properties"`
	Required   []ErrorRequired `json:"required"`
	Type       TypeElement     `json:"type"`
}
