package main

type CategoryType int

const (
	CategoryTypeUnknown = iota
	CategoryTypeDrug
	CategoryTypeService2
	CategoryTypeService3
	CategoryTypeMaterial
)

var categoryTypes = map[CategoryType]string{
	CategoryTypeDrug:     "药品",
	CategoryTypeService2: "治疗服务2",
	CategoryTypeService3: "治疗服务3",
	CategoryTypeMaterial: "医用材料",
}

func (v CategoryType) String() string {
	if desc, ok := categoryTypes[v]; ok {
		return desc
	}

	return "UNKNOWN"
}
