package v20260728

type InputRequiredResultProperties struct {
	Meta          AdditionalProperties `json:"_meta"`
	InputRequests AdditionalProperties `json:"inputRequests"`
	RequestState  Default              `json:"requestState"`
	ResultType    Cursor               `json:"resultType"`
}
