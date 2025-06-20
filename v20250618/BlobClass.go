package v20250618

type BlobClass struct {
	Description string      `json:"description"`
	Format      Required    `json:"format"`
	Type        TypeElement `json:"type"`
}
