package v20250618

type InitializeResultProperties struct {
	Meta            Meta        `json:"_meta"`
	Capabilities    EmptyResult `json:"capabilities"`
	Instructions    Cursor      `json:"instructions"`
	ProtocolVersion Cursor      `json:"protocolVersion"`
	ServerInfo      EmptyResult `json:"serverInfo"`
}
