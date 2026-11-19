package vdraft

type RootProperties struct {
	Meta AdditionalProperties `json:"_meta"`
	Name Cursor               `json:"name"`
	URI  BlobClass            `json:"uri"`
}
