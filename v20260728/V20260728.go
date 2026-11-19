// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    v20260728, err := UnmarshalV20260728(bytes)
//    bytes, err = v20260728.Marshal()

package v20260728

import "encoding/json"

func UnmarshalV20260728(data []byte) (V20260728, error) {
	var r V20260728
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *V20260728) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type V20260728 struct {
	Schema string `json:"$schema"`
	Defs   Defs   `json:"$defs"`
}
