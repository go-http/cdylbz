package main

//药品或者检查项目、材料
type Item struct {
	MedicalType MedicalType  `json:"aka063"`        //药品类别
	ReimburType ReimburType  `json:"aka065"`        //收费等级/限价：1甲2乙3丙
	Madein      MadeinType   `json:"ake004"`        //生产地类别: 1国内2国外
	Category    CategoryType `json:"ake003,string"` //编码类型，1药品、2/3诊疗服务、4医用材料

	UpdateReason   string `json:"aae013"` //??更新原因
	Code           string `json:"aaz231"` //药品编码/药品编码/材料编码
	GenericName    string `json:"aka061"` //通用名/项目名称/材料名称
	GenericName2   string `json:"bka003"` //?也是名称
	Produce        string `json:"bka005"` //生产单位
	BusinessName   string `json:"bka034"` //商品名
	Spec           string `json:"aka074"` //规格
	MedicamentType string `json:"bka004"` //剂型
	Limit          string `json:"bke047"` //限制使用范围
	Remark         string `json:"bka008"` //项目内涵

	//非丙类项目限价
	PriceTopA    string `json:"yka227"` //三甲医院限价
	PriceTopB    string `json:"yka228"` //三乙医院限价
	PriceSecondA string `json:"yka229"` //二甲医院限价
	PriceSecondB string `json:"yka230"` //二乙医院限价
	PriceClassC  string `json:"yka231"` //一级以下医院限价
	PriceMisc    string `json:"yka345"` //未定级医院限价

	//未知字段
	Aka062 string `json:"aka062"`
	Aka163 string `json:"aka163"`
	Bka009 string `json:"bka009"`
}
