package vdraft

type StickyProperties struct {
	Cancel   Metadata       `json:"cancel"`
	List     Metadata       `json:"list"`
	Requests PurpleRequests `json:"requests"`
}
