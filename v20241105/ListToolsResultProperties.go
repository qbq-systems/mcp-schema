package v20241105

type ListToolsResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
	Tools      Messages   `json:"tools"`
}
