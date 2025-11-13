package vdraft

type ElicitRequestURLParamsProperties struct {
	Meta          Meta      `json:"_meta"`
	ElicitationID Cursor    `json:"elicitationId"`
	Message       Cursor    `json:"message"`
	Mode          Mode      `json:"mode"`
	URL           BlobClass `json:"url"`
}
