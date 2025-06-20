package v20250326

type MischievousProperties struct {
	Message       Cursor           `json:"message"`
	Progress      Cursor           `json:"progress"`
	ProgressToken AnnotationsClass `json:"progressToken"`
	Total         Cursor           `json:"total"`
}
