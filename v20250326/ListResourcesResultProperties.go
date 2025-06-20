package v20250326

type ListResourcesResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
	Resources  Messages   `json:"resources"`
}
