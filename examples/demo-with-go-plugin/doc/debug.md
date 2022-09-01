### 调试源码

- `build examples/demo-with-go-plugin/plugin/debugtalk.go -o examples/demo-with-go-plugin`：编译二进制文件
- `run examples/demo-with-go-plugin/testcases/ref_testcase.yml --gen-html-report`：运行测试用例

### 修改httprunner之后，编译可执行文件

- 在项目的根目录下执行：`go install hrp/cmd/cli/main.go`