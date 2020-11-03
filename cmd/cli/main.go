package main

import (
	"flag"
	"fmt"

	"github.com/go-http/cdylbz"
)

var queryDrug, queryService, queryMaterial string

func init() {
	flag.StringVar(&queryDrug, "d", "", "药品名")
	flag.StringVar(&queryService, "s", "", "医疗服务名")
	flag.StringVar(&queryMaterial, "m", "", "医用材料名")
}

func main() {
	flag.Parse()
	if queryDrug == "" && queryService == "" && queryMaterial == "" {
		flag.Usage()
		return
	}

	var items []cdylbz.Item
	if queryDrug != "" {
		drugItems, err := cdylbz.QueryDrug(queryDrug)
		if err != nil {
			fmt.Println(err)
		}
		items = append(items, drugItems...)
	}

	if queryService != "" {
		serviceItems, err := cdylbz.QueryService(queryService)
		if err != nil {
			fmt.Println(err)
		}
		items = append(items, serviceItems...)
	}

	if queryMaterial != "" {
		materialItems, err := cdylbz.QueryMaterial(queryMaterial)
		if err != nil {
			fmt.Println(err)
		}
		items = append(items, materialItems...)
	}

	// 输出查询结果
	for _, item := range items {
		fmt.Println(item)
	}
}
