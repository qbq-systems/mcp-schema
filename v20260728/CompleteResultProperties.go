package v20260728

type CompleteResultProperties struct {
	Meta       AdditionalProperties `json:"_meta"`
	Completion Completion           `json:"completion"`
	ResultType Cursor               `json:"resultType"`
}
