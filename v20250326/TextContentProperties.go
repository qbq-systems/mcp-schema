package v20250326

type TextContentProperties struct {
	Annotations AnnotationsClass `json:"annotations"`
	Text        Cursor           `json:"text"`
	Type        TypeClass        `json:"type"`
}
