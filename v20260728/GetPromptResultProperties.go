package v20260728

type GetPromptResultProperties struct {
	Meta        AdditionalProperties `json:"_meta"`
	Description Cursor               `json:"description"`
	Messages    JSONArray            `json:"messages"`
	ResultType  Cursor               `json:"resultType"`
}
