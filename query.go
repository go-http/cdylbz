package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const ServerUrl = "http://cdsylbz.chengdu.gov.cn/cdsi_wxgzh/fwcUser/serchInfo"

var medicalTypes = map[string]string{
	"01": "西药",
	"02": "中成药",
	"03": "中草药",
	"04": "常规检查",
	"05": "CT ",
	"06": "核磁 ",
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
	return "Unknown"
}

type ReimburType string

var applyTypes = map[string]string{
	"1": "社保甲类", "2": "社保乙类", "3": "自费丙类",
}

func (v ReimburType) String() string {
	if desc, ok := applyTypes[string(v)]; ok {
		return desc
	}
	return string(v)
}

type MadeInType string

var madeInTypes = map[string]string{
	"1": "国内", "2": "国外",
}

func (v MadeInType) String() string {
	if desc, ok := madeInTypes[string(v)]; ok {
		return desc
	}
	return string(v)
}

type CategoryType string

var categoryTypes = map[string]string{
	"1": "药品", "2": "治疗服务", "3": "治疗服务3", "4": "医用材料",
}

func (v CategoryType) String() string {
	if desc, ok := categoryTypes[string(v)]; ok {
		return desc
	}
	return string(v)
}

type Product struct {
	MedicalType MedicalType  `json:"aka063"` //药品类别
	ReimburType ReimburType  `json:"aka065"` //收费等级/限价：1甲2乙3丙
	MadeIn      MadeInType   `json:"ake004"` //生产地类别: 1国内2国外
	Category    CategoryType `json:"ake003"` //编码类型，1药品、2/3诊疗服务、4医用材料

	UpdateReason   string `json:"aae013"` //??更新原因
	Code           string `json:"aaz231"` //药品编码/药品编码/材料编码
	Name           string `json:"aka061"` //通用名/项目名称/材料名称
	Name2          string `json:"bka003"` //?也是名称
	Produce        string `json:"bka005"` //生产单位
	BusinessName   string `json:"bka034"` //商品名
	Spec           string `json:"aka074"` //规格
	MedicamentType string `json:"bka004"` //剂型
	Limit          string `json:"bke047"` //限制使用范围
	Remark         string `json:"bka008"` //项目内涵

	//非丙类项目限价
	Price3J   string `json:"yka227"` //三甲医院限价
	Price3Y   string `json:"yka228"` //三乙医院限价
	Price2J   string `json:"yka229"` //二甲医院限价
	Price2Y   string `json:"yka230"` //二乙医院限价
	Price1    string `json:"yka231"` //一级以下医院限价
	PriceMisc string `json:"yka345"` //未定级医院限价

	//未知字段
	Aka062 string `json:"aka062"`
	Aka163 string `json:"aka163"`
	Bka009 string `json:"bka009"`
}

func QueryMedical(keyword string) ([]Product, error) {
	return Query("1", keyword)
}
func QueryService(keyword string) ([]Product, error) {
	return Query("2", keyword)
}
func QueryMaterial(keyword string) ([]Product, error) {
	return Query("4", keyword)
}

//支持名称、编码查询
func Query(category, keyword string) ([]Product, error) {

	p := url.Values{
		"jybh":      {"106"}, //固定值
		"pagenum":   {"1"},
		"ake003":    {category}, //查询类别，1药品、2或3诊疗服务、4医用材料
		"condition": {keyword},  //类别
	}

	resp, err := http.Get(ServerUrl + "?" + p.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respInfo struct {
		OutBusinessContent struct {
			Aka190   string
			Akc040   string
			DataRow  json.RawMessage
			PageNum  int `json:",string"`
			TotalNum int `json:",string"`
		}
		Outidentity struct {
			ReturnId  int `json:",string"`
			ReturnMsg string
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	if respInfo.Outidentity.ReturnId != 0 {
		return nil, fmt.Errorf("[%d]%s", resp.StatusCode, respInfo.Outidentity.ReturnMsg)
	}

	//单个返回值结构比较奇葩
	if respInfo.OutBusinessContent.TotalNum == 1 {
		var singleData struct {
			Row Product
		}
		err = json.Unmarshal(respInfo.OutBusinessContent.DataRow, &singleData)
		if err != nil {
			return nil, err
		}
		return []Product{singleData.Row}, nil

	}

	//多个返回值
	products := make([]Product, 0, 10)
	err = json.Unmarshal(respInfo.OutBusinessContent.DataRow, &products)
	if err != nil {
		return nil, err
	}

	return products, nil

}
