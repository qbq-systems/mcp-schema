package v20250618

type ListToolsResultProperties struct {
	Meta       Meta     `json:"_meta"`
	NextCursor Cursor   `json:"nextCursor"`
	Tools      Messages `json:"tools"`
}
