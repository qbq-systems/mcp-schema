package vdraft

type FormatElement string

const (
	Byte        FormatElement = "byte"
	RequiredURI FormatElement = ": uri"
	URI         FormatElement = "uri"
	URITemplate FormatElement = "uri-template"
)
