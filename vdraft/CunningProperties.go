package vdraft

type CunningProperties struct {
	Capabilities    EmptyResult `json:"capabilities"`
	ClientInfo      EmptyResult `json:"clientInfo"`
	ProtocolVersion Cursor      `json:"protocolVersion"`
}
