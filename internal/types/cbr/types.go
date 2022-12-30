package cbr

import (
	"encoding/json"
	"encoding/xml"
)

type Fiat struct {
	Name   xml.Name `xml:"ValCurs"`
	Valute []*struct {
		CharCode string `xml:"CharCode" json:"char_code"`
		Nominal  int    `xml:"Nominal" json:"nominal"`
		Value    string `xml:"Value" json:"value"`
	} `xml:"Valute" json:"valute"`
}

func (f Fiat) Marshal() ([]byte, error) {
	return json.Marshal(f.Valute)
}
