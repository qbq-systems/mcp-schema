package v20250326

type ListToolsResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
	Tools      Messages   `json:"tools"`
}
