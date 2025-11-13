package vdraft

type ElicitRequestFormParamsProperties struct {
	Meta            Meta            `json:"_meta"`
	Message         Cursor          `json:"message"`
	Mode            Mode            `json:"mode"`
	RequestedSchema RequestedSchema `json:"requestedSchema"`
}
