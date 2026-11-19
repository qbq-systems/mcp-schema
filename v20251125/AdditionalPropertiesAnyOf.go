package v20251125

type AdditionalPropertiesAnyOf struct {
	Items *Default   `json:"items,omitempty"`
	Type  *TypeUnion `json:"type"`
}
