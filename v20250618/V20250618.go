// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    v20250618, err := UnmarshalV20250618(bytes)
//    bytes, err = v20250618.Marshal()

package v20250618

import "encoding/json"

func UnmarshalV20250618(data []byte) (V20250618, error) {
	var r V20250618
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *V20250618) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type V20250618 struct {
	Schema      string      `json:"$schema"`
	Definitions Definitions `json:"definitions"`
}
