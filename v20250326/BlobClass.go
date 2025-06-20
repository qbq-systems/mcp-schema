package v20250326

type BlobClass struct {
	Description string      `json:"description"`
	Format      Required    `json:"format"`
	Type        TypeElement `json:"type"`
}
