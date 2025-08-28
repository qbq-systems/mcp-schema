package vdraft

type BlobClass struct {
	Description string        `json:"description"`
	Format      FormatElement `json:"format"`
	Type        TypeElement   `json:"type"`
}
