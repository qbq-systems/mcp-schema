package vdraft

type SamplingMessageProperties struct {
	Meta    Meta          `json:"_meta"`
	Content PurpleContent `json:"content"`
	Role    EmptyResult   `json:"role"`
}
