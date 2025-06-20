// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    vdraft, err := UnmarshalVdraft(bytes)
//    bytes, err = vdraft.Marshal()

package vdraft

import "encoding/json"

func UnmarshalVdraft(data []byte) (Vdraft, error) {
	var r Vdraft
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Vdraft) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Vdraft struct {
	Schema      string      `json:"$schema"`
	Definitions Definitions `json:"definitions"`
}
