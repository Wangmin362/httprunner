package main

import (
	"gitcdteam.skyguardmis.com/bigdt/gokit/pkg/pathx"
	"github.com/httprunner/httprunner/v4/hrp"
	"net/http"
	"os"
	"path/filepath"
)

var hrpRunner = hrp.NewRunner(nil).SetFailfast(false).GenHTMLReport().SetRequestsLogOn()

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// 1、生成debugtalk.bin插件
	debugtalkbinPath, err := filepath.Abs("examples/sk-health-check/debugtalk.bin")
	if err != nil {
		panic(err)
	}
	exist, err := pathx.PathExists(debugtalkbinPath)
	if err != nil {
		panic(err)
	}
	if !exist {
		if err := hrp.BuildPlugin("examples/sk-health-check/plugin/debugtalk.go", "."); err != nil {
			panic(err)
		}
	}

	// 2、把endpoint地址写入到 env 文件当中
	if err := os.WriteFile(".env", []byte("endpoint=endpoint=cd-ucss-230.gatorcloud.skyguardmis.com"), 0644); err != nil {
		panic(err)
	}

	// 3、启动测试用例
	testcase := hrp.TestCasePath("examples/sk-health-check/testcases/sps.yml")
	if err := hrpRunner.Run(&testcase); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/test", HelloHandler)
	http.ListenAndServe(":8100", nil)
}
