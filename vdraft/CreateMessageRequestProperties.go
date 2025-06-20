package vdraft

type CreateMessageRequestProperties struct {
	Method MethodClass  `json:"method"`
	Params StickyParams `json:"params"`
}
