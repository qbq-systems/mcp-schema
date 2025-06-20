package v20250618

type PaginatedResultProperties struct {
	Meta       Meta   `json:"_meta"`
	NextCursor Cursor `json:"nextCursor"`
}
