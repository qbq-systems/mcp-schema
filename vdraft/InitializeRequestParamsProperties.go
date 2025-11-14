package vdraft

type InitializeRequestParamsProperties struct {
	Meta            GetTaskPayloadResult `json:"_meta"`
	Capabilities    EmptyResult          `json:"capabilities"`
	ClientInfo      EmptyResult          `json:"clientInfo"`
	ProtocolVersion Cursor               `json:"protocolVersion"`
	Task            AnnotationsElement   `json:"task"`
}
