package main

import (
	"html/template"
	"net/http"

	"github.com/athurg/cdylbz"
)

const html = `
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://unpkg.com/purecss@2.0.3/build/pure-min.css" crossorigin="anonymous">
		<style>
			*{text-align:center;}
			table{margin:0 auto;}
			.reimbur-1{color:green;}
			.reimbur-2{color:blue;}
			.reimbur-3{color:brown;}
		</style>
	</head>
	<body>
		<div class="pure-g">
			<div class="pure-u-1-24"></div>
			<div class="pure-u-22-24">
				<form class="pure-form" method="POST">
					<select name="querytype">
						<option value="drug">药品</option>
						<option value="service">医疗服务</option>
						<option value="material">医用材料</option>
					</select>
					<input name="keyword" value="{{.keyword}}"/>
					<button type="submit" class="pure-button pure-button-primary">查询</button>
				</form>

				{{with .items}}
				<table class="pure-table pure-table-odd">
					<thead>
						<tr>
							<th>药品编码</th>
							<th>药品类型</th>
							<th>通用名称</th>
							<th>商品名称</th>
							<th>报销类型</th>
							<th>规格</th>
							<th>剂型</th>
							<th>厂商</th>
						</tr>
					</thead>
					<tbody>
						{{range .}}
						<tr>
							<td>{{.Code}}</td>
							<td>{{.MedicalType}}</td>
							<td>{{.GenericName}}</td>
							<td>{{.BusinessName}}</td>
							<td class="reimbur-{{printf "%d" .ReimburType}}">{{.ReimburType}}</td>
							<td>{{.Spec}}</td>
							<td>{{.MedicamentType}}</td>
							<td>{{.Produce}}</td>
						</tr>
						{{end}}
					</tbody>
				</table>
				{{end}}
			</div>
			<div class="pure-u-1-24"></div>
		</div>
	</body>
</html>`

var tpl = template.Must(template.New("index").Parse(html))

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")

	templateData := map[string]interface{}{
		"keyword": keyword,
	}

	defer tpl.Execute(w, templateData)

	if keyword == "" {
		return
	}

	var err error
	var items []cdylbz.Item

	switch r.FormValue("querytype") {
	case "drug":
		items, err = cdylbz.QueryDrug(keyword)
	case "service":
		items, err = cdylbz.QueryService(keyword)
	case "material":
		items, err = cdylbz.QueryMaterial(keyword)
	}

	if err != nil {
		templateData["error"] = err
		return
	}

	templateData["items"] = items
}
