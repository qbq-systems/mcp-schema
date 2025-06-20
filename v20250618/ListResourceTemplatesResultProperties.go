package v20250618

type ListResourceTemplatesResultProperties struct {
	Meta              Meta     `json:"_meta"`
	NextCursor        Cursor   `json:"nextCursor"`
	ResourceTemplates Messages `json:"resourceTemplates"`
}
