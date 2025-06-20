package vdraft

type ErrorProperties struct {
	Code    Cursor     `json:"code"`
	Data    PurpleData `json:"data"`
	Message Cursor     `json:"message"`
}
