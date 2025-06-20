package vdraft

type RootProperties struct {
	Meta Meta      `json:"_meta"`
	Name Cursor    `json:"name"`
	URI  BlobClass `json:"uri"`
}
