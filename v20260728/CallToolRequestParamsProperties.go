package v20260728

type CallToolRequestParamsProperties struct {
	Meta           AdditionalProperties `json:"_meta"`
	Arguments      OutputSchemaClass    `json:"arguments"`
	InputResponses AdditionalProperties `json:"inputResponses"`
	Name           Cursor               `json:"name"`
	RequestState   Default              `json:"requestState"`
}
