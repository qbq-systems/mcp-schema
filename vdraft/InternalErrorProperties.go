package vdraft

type InternalErrorProperties struct {
	Code    FluffyCode        `json:"code"`
	Data    StructuredContent `json:"data"`
	Message Cursor            `json:"message"`
}
