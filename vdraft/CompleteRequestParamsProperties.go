package vdraft

type CompleteRequestParamsProperties struct {
	Meta     Meta               `json:"_meta"`
	Argument Argument           `json:"argument"`
	Context  Context            `json:"context"`
	Ref      ClientNotification `json:"ref"`
}
