package vdraft

type PaginatedResultProperties struct {
	Meta       Meta   `json:"_meta"`
	NextCursor Cursor `json:"nextCursor"`
}
