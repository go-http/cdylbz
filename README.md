# 成都医保查询工具

用于查询指定药品或者医疗服务、医用材料是否符合成都市医保报销标准。

附带了两个范例：

- [cmd/cli](./cmd/cli/main.go) 命令行工具，通过命令行查询
- [cmd/webserver](./cmd/webserver/main.go) 简单的WebServer，启动后可以通过访问 http://localhost:8080/ 查询。
