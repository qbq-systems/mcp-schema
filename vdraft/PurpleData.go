package vdraft

type PurpleData struct {
	Properties StickyProperties `json:"properties"`
	Required   []string         `json:"required"`
	Type       TypeElement      `json:"type"`
}
