package v20250618

type BraggadociousProperties struct {
	Message       Cursor           `json:"message"`
	Progress      Cursor           `json:"progress"`
	ProgressToken AnnotationsClass `json:"progressToken"`
	Total         Cursor           `json:"total"`
}
