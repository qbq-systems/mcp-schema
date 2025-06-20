package v20250326

type JsonrpcBatchRequest struct {
	Description *string             `json:"description,omitempty"`
	Items       *ClientNotification `json:"items,omitempty"`
	Type        *string             `json:"type,omitempty"`
	Ref         *string             `json:"$ref,omitempty"`
}
