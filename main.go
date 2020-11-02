package main

import (
	"fmt"
)

func main() {
	products, err := QueryMedical("86904011000038")
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range products {
		fmt.Printf("%+v", product)
	}

	products, err = QueryMedical("小儿消积")
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	products, err = QueryService("彩色多普勒")
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	products, err = QueryMaterial("支架")
	if err != nil {
		fmt.Println(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}
