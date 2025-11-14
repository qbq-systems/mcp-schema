package vdraft

type CreateMessageRequestParamsProperties struct {
	Meta             Meta             `json:"_meta"`
	IncludeContext   LoggingLevel     `json:"includeContext"`
	MaxTokens        Cursor           `json:"maxTokens"`
	Messages         Messages         `json:"messages"`
	Metadata         Metadata         `json:"metadata"`
	ModelPreferences AnnotationsClass `json:"modelPreferences"`
	StopSequences    StopSequences    `json:"stopSequences"`
	SystemPrompt     Cursor           `json:"systemPrompt"`
	Temperature      Default          `json:"temperature"`
	ToolChoice       AnnotationsClass `json:"toolChoice"`
	Tools            Audience         `json:"tools"`
}
