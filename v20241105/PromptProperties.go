package v20241105

type PromptProperties struct {
	Arguments   Hints  `json:"arguments"`
	Description Cursor `json:"description"`
	Name        Cursor `json:"name"`
}
