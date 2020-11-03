package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const ServerUrl = "http://cdsylbz.chengdu.gov.cn/cdsi_wxgzh/fwcUser/serchInfo"

func QueryDrug(keyword string) ([]Item, error) {
	return QueryPage(CategoryTypeDrug, keyword, 1)
}

func QueryService(keyword string) ([]Item, error) {
	return QueryPage(CategoryTypeService2, keyword, 1)
}

func QueryMaterial(keyword string) ([]Item, error) {
	return QueryPage(CategoryTypeMaterial, keyword, 1)
}

//支持名称、编码查询
func QueryPage(category CategoryType, keyword string, pageNum int) ([]Item, error) {

	p := url.Values{
		"jybh":      {"106"}, //固定值
		"pagenum":   {fmt.Sprintf("%d", pageNum)},
		"ake003":    {fmt.Sprintf("%d", category)},
		"condition": {keyword},
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
			Row Item
		}
		err = json.Unmarshal(respInfo.OutBusinessContent.DataRow, &singleData)
		if err != nil {
			return nil, err
		}
		return []Item{singleData.Row}, nil

	}

	//多个返回值
	items := make([]Item, 0, 10)
	err = json.Unmarshal(respInfo.OutBusinessContent.DataRow, &items)
	if err != nil {
		return nil, err
	}

	return items, nil

}
