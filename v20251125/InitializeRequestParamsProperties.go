package v20251125

type InitializeRequestParamsProperties struct {
	Meta            GetTaskPayloadResult `json:"_meta"`
	Capabilities    EmptyResult          `json:"capabilities"`
	ClientInfo      EmptyResult          `json:"clientInfo"`
	ProtocolVersion Cursor               `json:"protocolVersion"`
}
