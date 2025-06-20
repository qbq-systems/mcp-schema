// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    v20241105, err := UnmarshalV20241105(bytes)
//    bytes, err = v20241105.Marshal()

package v20241105

import "encoding/json"

func UnmarshalV20241105(data []byte) (V20241105, error) {
	var r V20241105
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *V20241105) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type V20241105 struct {
	Schema      string      `json:"$schema"`
	Definitions Definitions `json:"definitions"`
}
