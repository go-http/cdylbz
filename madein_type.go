package main

type MadeinType string

var madeinTypes = map[string]string{
	"1": "国内",
	"2": "国外",
}

func (v MadeinType) String() string {
	if desc, ok := madeinTypes[string(v)]; ok {
		return desc
	}

	return string(v)
}
