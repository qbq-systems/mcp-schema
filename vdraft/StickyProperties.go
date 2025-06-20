package vdraft

type StickyProperties struct {
	Argument Argument           `json:"argument"`
	Context  Context            `json:"context"`
	Ref      ClientNotification `json:"ref"`
}
