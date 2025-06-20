package vdraft

type InitializedNotificationProperties struct {
	Method MethodClass `json:"method"`
	Params Result      `json:"params"`
}
