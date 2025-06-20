package v20241105

type Blob struct {
	Description string      `json:"description"`
	Format      Required    `json:"format"`
	Type        TypeElement `json:"type"`
}
