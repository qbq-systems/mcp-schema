package vdraft

type ElicitRequestParamsProperties struct {
	Meta            Meta            `json:"_meta"`
	Message         Cursor          `json:"message"`
	RequestedSchema RequestedSchema `json:"requestedSchema"`
}
