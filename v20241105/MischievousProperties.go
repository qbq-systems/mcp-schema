package v20241105

type MischievousProperties struct {
	Progress      Cursor    `json:"progress"`
	ProgressToken RequestID `json:"progressToken"`
	Total         Cursor    `json:"total"`
}
