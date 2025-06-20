package v20241105

type TextContentProperties struct {
	Annotations Annotations `json:"annotations"`
	Text        Cursor      `json:"text"`
	Type        Method      `json:"type"`
}
