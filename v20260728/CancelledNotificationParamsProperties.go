package v20260728

type CancelledNotificationParamsProperties struct {
	Meta      AdditionalProperties `json:"_meta"`
	Reason    Cursor               `json:"reason"`
	RequestID ClientResult         `json:"requestId"`
}
