package v20241105

type ListResourceTemplatesResultProperties struct {
	Meta              PurpleMeta `json:"_meta"`
	NextCursor        Cursor     `json:"nextCursor"`
	ResourceTemplates Messages   `json:"resourceTemplates"`
}
