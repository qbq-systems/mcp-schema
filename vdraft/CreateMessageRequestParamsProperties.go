package vdraft

type CreateMessageRequestParamsProperties struct {
	Meta             AdditionalProperties `json:"_meta"`
	IncludeContext   LoggingLevel         `json:"includeContext"`
	MaxTokens        Cursor               `json:"maxTokens"`
	Messages         JSONArray            `json:"messages"`
	Metadata         ClientResult         `json:"metadata"`
	ModelPreferences ClientResult         `json:"modelPreferences"`
	StopSequences    StopSequences        `json:"stopSequences"`
	SystemPrompt     Cursor               `json:"systemPrompt"`
	Temperature      Default              `json:"temperature"`
	ToolChoice       ClientResult         `json:"toolChoice"`
	Tools            Audience             `json:"tools"`
}
