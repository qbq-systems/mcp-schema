package v20250326

type InitializeResultProperties struct {
	Meta            PurpleMeta  `json:"_meta"`
	Capabilities    EmptyResult `json:"capabilities"`
	Instructions    Cursor      `json:"instructions"`
	ProtocolVersion Cursor      `json:"protocolVersion"`
	ServerInfo      EmptyResult `json:"serverInfo"`
}
