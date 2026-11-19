package v20251125

type InitializeResultProperties struct {
	Meta            GetTaskPayloadResult `json:"_meta"`
	Capabilities    EmptyResult          `json:"capabilities"`
	Instructions    Cursor               `json:"instructions"`
	ProtocolVersion Cursor               `json:"protocolVersion"`
	ServerInfo      EmptyResult          `json:"serverInfo"`
}
