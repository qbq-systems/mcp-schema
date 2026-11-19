package vdraft

type GetPromptRequestParamsProperties struct {
	Meta           AdditionalProperties `json:"_meta"`
	Arguments      PurpleArguments      `json:"arguments"`
	InputResponses AdditionalProperties `json:"inputResponses"`
	Name           Cursor               `json:"name"`
	RequestState   Default              `json:"requestState"`
}
