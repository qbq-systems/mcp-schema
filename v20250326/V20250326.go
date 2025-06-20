// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    v20250326, err := UnmarshalV20250326(bytes)
//    bytes, err = v20250326.Marshal()

package v20250326

import "encoding/json"

func UnmarshalV20250326(data []byte) (V20250326, error) {
	var r V20250326
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *V20250326) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type V20250326 struct {
	Schema      string      `json:"$schema"`
	Definitions Definitions `json:"definitions"`
}
