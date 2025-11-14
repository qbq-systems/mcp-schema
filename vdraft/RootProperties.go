package vdraft

type RootProperties struct {
	Meta GetTaskPayloadResult `json:"_meta"`
	Name Cursor               `json:"name"`
	URI  BlobClass            `json:"uri"`
}
