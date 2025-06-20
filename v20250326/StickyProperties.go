package v20250326

type StickyProperties struct {
	IncludeContext   LoggingLevel     `json:"includeContext"`
	MaxTokens        Cursor           `json:"maxTokens"`
	Messages         Messages         `json:"messages"`
	Metadata         Sampling         `json:"metadata"`
	ModelPreferences AnnotationsClass `json:"modelPreferences"`
	StopSequences    StopSequences    `json:"stopSequences"`
	SystemPrompt     Cursor           `json:"systemPrompt"`
	Temperature      Name             `json:"temperature"`
}
