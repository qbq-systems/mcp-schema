package v20251125

type BlobClass struct {
	Description string        `json:"description"`
	Format      FormatElement `json:"format"`
	Type        TypeElement   `json:"type"`
}
