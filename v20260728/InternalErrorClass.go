package v20260728

type InternalErrorClass struct {
	Description string                  `json:"description"`
	Properties  InternalErrorProperties `json:"properties"`
	Required    []ErrorRequired         `json:"required"`
	Type        TypeElement             `json:"type"`
}
