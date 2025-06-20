package vdraft

type ListResourceTemplatesResultProperties struct {
	Meta              Meta     `json:"_meta"`
	NextCursor        Cursor   `json:"nextCursor"`
	ResourceTemplates Messages `json:"resourceTemplates"`
}
