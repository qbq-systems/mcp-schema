package vdraft

type InitializeRequestParamsProperties struct {
	Meta            Meta        `json:"_meta"`
	Capabilities    EmptyResult `json:"capabilities"`
	ClientInfo      EmptyResult `json:"clientInfo"`
	ProtocolVersion Cursor      `json:"protocolVersion"`
}
