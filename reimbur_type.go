package cdylbz

type ReimburType string

var applyTypes = map[string]string{
	"1": "社保甲类",
	"2": "社保乙类",
	"3": "自费丙类",
}

func (v ReimburType) String() string {
	if desc, ok := applyTypes[string(v)]; ok {
		return desc
	}

	return string(v)
}
