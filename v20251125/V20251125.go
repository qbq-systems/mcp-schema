// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    v20251125, err := UnmarshalV20251125(bytes)
//    bytes, err = v20251125.Marshal()

package v20251125

import "encoding/json"

func UnmarshalV20251125(data []byte) (V20251125, error) {
	var r V20251125
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *V20251125) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type V20251125 struct {
	Schema string `json:"$schema"`
	Defs   Defs   `json:"$defs"`
}
