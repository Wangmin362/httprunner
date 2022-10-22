### 调试源码

- `build examples/sk-health-check/plugin/debugtalk.go -o examples/sk-health-check`：编译二进制文件
- `run examples/sk-health-check/testcases/sps.yml --gen-html-report`：运行测试用例
- `run examples/sk-health-check/testcases/ucwi.yml --gen-html-report`：运行测试用例

- `boom examples/sk-health-check/testcases/ucwi.yml --spawn-count 1000 --spawn-rate 100`：运行压力测试

### 修改httprunner之后，编译可执行文件

- 在项目的根目录下执行：`go install hrp/cmd/cli/main.go`
