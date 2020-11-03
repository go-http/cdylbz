package main

var medicalTypes = map[string]string{
	"01": "西药",
	"02": "中成药",
	"03": "中草药",
	"04": "常规检查",
	"05": "CT",
	"06": "核磁",
	"07": "B超",
	"08": "治疗费",
	"09": "化验费",
	"10": "手术费",
	"11": "输氧费",
	"12": "放射费",
	"13": "输血费",
	"14": "注射费",
	"15": "透析费",
	"16": "化疗费",
	"17": "床位费",
	"18": "材料费",
	"19": "护理费",
	"99": "其他费用",
}

type MedicalType string

func (v MedicalType) String() string {
	if desc, ok := medicalTypes[string(v)]; ok {
		return desc
	}

	return string(v)
}
