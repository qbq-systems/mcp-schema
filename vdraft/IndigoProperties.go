package vdraft

type IndigoProperties struct {
	IncludeContext   LoggingLevel     `json:"includeContext"`
	MaxTokens        Cursor           `json:"maxTokens"`
	Messages         Messages         `json:"messages"`
	Metadata         Elicitation      `json:"metadata"`
	ModelPreferences AnnotationsClass `json:"modelPreferences"`
	StopSequences    Enum             `json:"stopSequences"`
	SystemPrompt     Cursor           `json:"systemPrompt"`
	Temperature      Default          `json:"temperature"`
}
