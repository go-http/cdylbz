package cdylbz

type ReimburType int

const (
	ReimburTypeUnknow ReimburType = iota
	ReimburTypeA
	ReimburTypeB
	ReimburTypeC
)

var reimburTypes = map[ReimburType]string{
	ReimburTypeA: "社保甲类",
	ReimburTypeB: "社保乙类",
	ReimburTypeC: "自费丙类",
}

func (v ReimburType) String() string {
	if desc, ok := reimburTypes[v]; ok {
		return desc
	}

	return "未知报销类型"
}
