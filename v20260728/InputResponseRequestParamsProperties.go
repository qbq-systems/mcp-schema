package v20260728

type InputResponseRequestParamsProperties struct {
	Meta           AdditionalProperties `json:"_meta"`
	InputResponses AdditionalProperties `json:"inputResponses"`
	RequestState   Default              `json:"requestState"`
	URI            *BlobClass           `json:"uri,omitempty"`
}
