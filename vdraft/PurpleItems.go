package vdraft

type PurpleItems struct {
	Description string              `json:"description"`
	Properties  HilariousProperties `json:"properties"`
	Required    []string            `json:"required"`
	Type        TypeElement         `json:"type"`
}
