package v20260728

type MissingRequiredClientCapabilityError struct {
	Description string                                         `json:"description"`
	Properties  MissingRequiredClientCapabilityErrorProperties `json:"properties"`
	Required    []string                                       `json:"required"`
	Type        TypeElement                                    `json:"type"`
}
