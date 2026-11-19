package vdraft

type MagentaProperties struct {
	Requested Cursor            `json:"requested"`
	Supported SupportedVersions `json:"supported"`
}
